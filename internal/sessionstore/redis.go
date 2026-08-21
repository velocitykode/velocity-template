// Package sessionstore implements velocity/auth.ServerSessionStore backed by
// the framework's cache.Manager. The platform's default cache driver is
// Redis (see CACHE_DRIVER in .env), so server-side session
// records survive platform restarts and stay coherent across instances.
//
// The framework's SessionScheme owns the cookie session and writes one
// StoredSession per Login into this store; this package owns the
// administrative view (list active sessions, revoke a single session,
// sign out everywhere). Manager.SetServerSessionStore propagates the
// store to the scheme via the auth.ServerSessionStoreReceiver interface,
// so this package never needs to be invoked by the app's request path.
package sessionstore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/cache"
)

// Key namespaces used inside the shared cache. The meta key holds the JSON
// StoredSession; the user index holds a JSON array of session IDs owned by
// a single user. Both share the session: prefix so SCAN session:* picks up
// every entry for ops-side eviction.
const (
	metaKeyPrefix = "session:meta:"
	userKeyPrefix = "session:user:"
)

// userIndexSlack is added to the per-user index TTL so the list of session
// ids outlives the longest meta key it points at, absorbing clock skew
// between the two writes.
const userIndexSlack = 1 * time.Minute

// ErrNilCache is returned when the constructor is called without a cache
// manager. Surfaces at boot so misconfigured modules fail fast.
var ErrNilCache = errors.New("sessionstore: cache manager is nil")

// Store is an auth.ServerSessionStore backed by velocity/cache. The "Redis"
// name reflects the production driver; the same struct works against the
// memory driver in tests because everything goes through cache.Manager.
//
// Performance note: cache.Manager exposes Get / Put / Forget but no native
// set ops, so DeleteAllForUser and ListForUser do a JSON-array round-trip
// on the per-user index key. For a single user with N concurrent sessions
// this is O(N) Redis ops + one JSON encode/decode. Real users carry single
// digits of active sessions, so the simpler shape beats reaching past the
// CacheManager interface for SADD/SMEMBERS.
type Store struct {
	cache cache.CacheManager
}

// Compile-time guarantee that Store satisfies the framework interface.
var _ auth.ServerSessionStore = (*Store)(nil)

// New builds a Store. ErrNilCache is returned when cm is nil so a missing
// cache surfaces during AppModule.Start rather than silently producing a
// store whose writes vanish.
func New(cm cache.CacheManager) (*Store, error) {
	if cm == nil {
		return nil, ErrNilCache
	}
	return &Store{cache: cm}, nil
}

func metaKey(id string) string { return metaKeyPrefix + id }
func userKey(id string) string { return userKeyPrefix + id }

// storedSessionDTO is the JSON shape persisted under session:meta:<id>.
// Mirrors auth.StoredSession but uses explicit json tags so the on-disk
// format is stable across framework refactors.
type storedSessionDTO struct {
	ID         string         `json:"id"`
	UserID     string         `json:"user_id"`
	Data       map[string]any `json:"data,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	LastSeenAt time.Time      `json:"last_seen_at"`
	ExpiresAt  time.Time      `json:"expires_at"`
	IPAddress  string         `json:"ip_address,omitempty"`
	UserAgent  string         `json:"user_agent,omitempty"`
}

func toDTO(s *auth.StoredSession) storedSessionDTO {
	return storedSessionDTO{
		ID:         s.ID,
		UserID:     s.UserID,
		Data:       s.Data,
		CreatedAt:  s.CreatedAt,
		LastSeenAt: s.LastSeenAt,
		ExpiresAt:  s.ExpiresAt,
		IPAddress:  s.IPAddress,
		UserAgent:  s.UserAgent,
	}
}

func (d storedSessionDTO) toStored() *auth.StoredSession {
	return &auth.StoredSession{
		ID:         d.ID,
		UserID:     d.UserID,
		Data:       d.Data,
		CreatedAt:  d.CreatedAt,
		LastSeenAt: d.LastSeenAt,
		ExpiresAt:  d.ExpiresAt,
		IPAddress:  d.IPAddress,
		UserAgent:  d.UserAgent,
	}
}

// Get returns the StoredSession for id. Missing keys map to ErrSessionNotFound;
// expired keys are deleted and reported as ErrSessionExpired so the framework
// contract matches the in-memory driver bit-for-bit.
func (s *Store) Get(ctx context.Context, id string) (*auth.StoredSession, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if id == "" {
		return nil, auth.ErrSessionNotFound
	}
	raw, ok := s.cache.GetWithContext(ctx, metaKey(id))
	if !ok {
		return nil, auth.ErrSessionNotFound
	}
	dto, err := decodeDTO(raw)
	if err != nil {
		return nil, fmt.Errorf("sessionstore: decode session %s: %w", id, err)
	}
	if !dto.ExpiresAt.IsZero() && time.Now().After(dto.ExpiresAt) {
		// Best-effort cleanup; ignore the error since the meta key may have
		// already been TTL-evicted out from under us.
		_ = s.cache.ForgetWithContext(ctx, metaKey(id))
		s.removeFromUserIndex(ctx, dto.UserID, id)
		return nil, auth.ErrSessionExpired
	}
	return dto.toStored(), nil
}

// Put creates or replaces a session record. ID and UserID are required.
// LastSeenAt is overwritten with time.Now() to match the contract; the
// caller's CreatedAt is preserved. The meta key TTL tracks ExpiresAt so a
// session evicted by Redis disappears from ListForUser too.
func (s *Store) Put(ctx context.Context, sess *auth.StoredSession) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if sess == nil {
		return errors.New("sessionstore: nil session")
	}
	if sess.ID == "" {
		return errors.New("sessionstore: empty session id")
	}
	if sess.UserID == "" {
		return errors.New("sessionstore: empty user id")
	}

	now := time.Now()
	dto := toDTO(sess)
	if dto.CreatedAt.IsZero() {
		dto.CreatedAt = now
	}
	dto.LastSeenAt = now

	// If the same id was previously bound to a different user, scrub the
	// stale entry from the old user's index so ListForUser stays consistent.
	// Read the prior user BEFORE overwriting the meta key.
	if previous, err := s.peekPriorUser(ctx, dto.ID, dto.UserID); err == nil && previous != "" {
		s.removeFromUserIndex(ctx, previous, dto.ID)
	}

	ttl := sessionTTL(dto.ExpiresAt, now)
	if err := s.cache.PutWithContext(ctx, metaKey(dto.ID), dto, ttl); err != nil {
		return fmt.Errorf("sessionstore: put meta: %w", err)
	}

	if err := s.addToUserIndex(ctx, dto.UserID, dto.ID, ttl); err != nil {
		// Roll back the meta write so we don't leak orphan rows the user
		// can never list or revoke.
		_ = s.cache.ForgetWithContext(ctx, metaKey(dto.ID))
		return fmt.Errorf("sessionstore: index put: %w", err)
	}
	return nil
}

// Delete removes a single session. Idempotent: missing ids return nil.
// Also removes the id from its owning user's index; failures are tolerated
// because a stale index entry is just a no-op on the next ListForUser
// (the meta lookup misses and the entry is skipped).
func (s *Store) Delete(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if id == "" {
		return nil
	}
	// Look up the user so we can scrub the index. Missing meta is fine;
	// an explicit Forget afterwards is still cheap.
	var userID string
	if raw, ok := s.cache.GetWithContext(ctx, metaKey(id)); ok {
		if dto, err := decodeDTO(raw); err == nil {
			userID = dto.UserID
		}
	}
	if err := s.cache.ForgetWithContext(ctx, metaKey(id)); err != nil {
		return fmt.Errorf("sessionstore: forget meta: %w", err)
	}
	if userID != "" {
		s.removeFromUserIndex(ctx, userID, id)
	}
	return nil
}

// DeleteAllForUser removes every session record belonging to userID and
// clears the per-user index. Returns nil when the user has no sessions.
func (s *Store) DeleteAllForUser(ctx context.Context, userID string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if userID == "" {
		return nil
	}
	ids, err := s.loadUserIndex(ctx, userID)
	if err != nil {
		return fmt.Errorf("sessionstore: load index: %w", err)
	}
	for _, id := range ids {
		if err := s.cache.ForgetWithContext(ctx, metaKey(id)); err != nil {
			return fmt.Errorf("sessionstore: forget meta %s: %w", id, err)
		}
	}
	if err := s.cache.ForgetWithContext(ctx, userKey(userID)); err != nil {
		return fmt.Errorf("sessionstore: forget index: %w", err)
	}
	return nil
}

// ListForUser returns SessionMeta for every non-expired session belonging
// to userID. Stale index entries (meta key gone or expired) are reaped on
// the way out so the index stays bounded. The Data field is intentionally
// omitted: administrative listings must not leak per-session payloads.
func (s *Store) ListForUser(ctx context.Context, userID string) ([]*auth.SessionMeta, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if userID == "" {
		return nil, nil
	}
	ids, err := s.loadUserIndex(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("sessionstore: load index: %w", err)
	}

	now := time.Now()
	out := make([]*auth.SessionMeta, 0, len(ids))
	stale := make([]string, 0)
	for _, id := range ids {
		raw, ok := s.cache.GetWithContext(ctx, metaKey(id))
		if !ok {
			stale = append(stale, id)
			continue
		}
		dto, err := decodeDTO(raw)
		if err != nil {
			stale = append(stale, id)
			continue
		}
		if !dto.ExpiresAt.IsZero() && now.After(dto.ExpiresAt) {
			_ = s.cache.ForgetWithContext(ctx, metaKey(id))
			stale = append(stale, id)
			continue
		}
		out = append(out, dto.toStored().ToMeta())
	}
	if len(stale) > 0 {
		s.scrubFromUserIndex(ctx, userID, stale)
	}
	return out, nil
}

// addToUserIndex appends id to the per-user index, deduplicating in case
// the same id is rewritten (e.g. last-seen refresh). TTL is bumped on
// every write so the index outlives the longest active meta key.
func (s *Store) addToUserIndex(ctx context.Context, userID, id string, ttl time.Duration) error {
	ids, err := s.loadUserIndex(ctx, userID)
	if err != nil {
		return err
	}
	indexTTL := ttl + userIndexSlack
	for _, existing := range ids {
		if existing == id {
			return s.writeUserIndex(ctx, userID, ids, indexTTL)
		}
	}
	ids = append(ids, id)
	return s.writeUserIndex(ctx, userID, ids, indexTTL)
}

// removeFromUserIndex removes a single id from the per-user index. Best
// effort: failures are swallowed because a stale index entry is harmless
// (ListForUser skips it) and the caller is on a delete path that already
// succeeded for the meta key.
func (s *Store) removeFromUserIndex(ctx context.Context, userID, id string) {
	if userID == "" || id == "" {
		return
	}
	ids, err := s.loadUserIndex(ctx, userID)
	if err != nil {
		return
	}
	filtered := ids[:0]
	for _, existing := range ids {
		if existing != id {
			filtered = append(filtered, existing)
		}
	}
	if len(filtered) == len(ids) {
		return // nothing to remove
	}
	if len(filtered) == 0 {
		_ = s.cache.ForgetWithContext(ctx, userKey(userID))
		return
	}
	// Reuse a long TTL on the index; we don't know the longest meta TTL
	// without re-reading every entry, and the meta keys are the source of
	// truth so a slightly long index TTL only delays self-cleanup.
	_ = s.writeUserIndex(ctx, userID, filtered, 7*24*time.Hour+userIndexSlack)
}

// scrubFromUserIndex removes a batch of ids from the per-user index in a
// single round-trip. Used by ListForUser to reap stale entries.
func (s *Store) scrubFromUserIndex(ctx context.Context, userID string, stale []string) {
	if userID == "" || len(stale) == 0 {
		return
	}
	ids, err := s.loadUserIndex(ctx, userID)
	if err != nil {
		return
	}
	staleSet := make(map[string]struct{}, len(stale))
	for _, id := range stale {
		staleSet[id] = struct{}{}
	}
	filtered := ids[:0]
	for _, existing := range ids {
		if _, drop := staleSet[existing]; !drop {
			filtered = append(filtered, existing)
		}
	}
	if len(filtered) == 0 {
		_ = s.cache.ForgetWithContext(ctx, userKey(userID))
		return
	}
	_ = s.writeUserIndex(ctx, userID, filtered, 7*24*time.Hour+userIndexSlack)
}

// loadUserIndex decodes the per-user JSON-array index. Missing keys are
// reported as nil + nil error so callers can append-or-create.
func (s *Store) loadUserIndex(ctx context.Context, userID string) ([]string, error) {
	raw, ok := s.cache.GetWithContext(ctx, userKey(userID))
	if !ok {
		return nil, nil
	}
	switch v := raw.(type) {
	case []string:
		return v, nil
	case []any:
		out := make([]string, 0, len(v))
		for _, e := range v {
			if s, ok := e.(string); ok {
				out = append(out, s)
			}
		}
		return out, nil
	default:
		// Unknown shape: re-marshal and JSON-decode so we accept any
		// driver's loose typing of slice-shaped values.
		buf, err := json.Marshal(raw)
		if err != nil {
			return nil, err
		}
		var ids []string
		if err := json.Unmarshal(buf, &ids); err != nil {
			return nil, err
		}
		return ids, nil
	}
}

// writeUserIndex persists the per-user id list with the given TTL.
func (s *Store) writeUserIndex(ctx context.Context, userID string, ids []string, ttl time.Duration) error {
	return s.cache.PutWithContext(ctx, userKey(userID), ids, ttl)
}

// peekPriorUser returns the user id currently bound to sessionID when it
// differs from incoming. Empty when no prior record exists or it already
// belonged to incoming. Used by Put to keep the per-user index consistent
// when a session id is reassigned.
func (s *Store) peekPriorUser(ctx context.Context, sessionID, incoming string) (string, error) {
	raw, ok := s.cache.GetWithContext(ctx, metaKey(sessionID))
	if !ok {
		return "", nil
	}
	dto, err := decodeDTO(raw)
	if err != nil {
		return "", err
	}
	if dto.UserID == incoming {
		return "", nil
	}
	return dto.UserID, nil
}

// decodeDTO normalises whatever shape the cache driver returned (Redis
// hands back a map[string]any, an in-process driver may hand back the
// original struct) into a typed storedSessionDTO via JSON round-trip.
func decodeDTO(raw any) (storedSessionDTO, error) {
	if dto, ok := raw.(storedSessionDTO); ok {
		return dto, nil
	}
	buf, err := json.Marshal(raw)
	if err != nil {
		return storedSessionDTO{}, err
	}
	var dto storedSessionDTO
	if err := json.Unmarshal(buf, &dto); err != nil {
		return storedSessionDTO{}, err
	}
	return dto, nil
}

// sessionTTL returns the cache-key TTL for a session: ExpiresAt - now,
// clamped to a sensible minimum so a near-expired session still gets
// written (test paths frequently set ExpiresAt = now + few ms). When
// ExpiresAt is the zero value we treat it as "no expiry" by returning 0,
// which cache.Manager interprets as "store forever" via Put.
func sessionTTL(expiresAt, now time.Time) time.Duration {
	if expiresAt.IsZero() {
		return 0
	}
	d := expiresAt.Sub(now)
	if d < time.Second {
		return time.Second
	}
	return d
}
