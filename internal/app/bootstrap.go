package app

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"{{MODULE_NAME}}/config"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/auth/drivers/guards"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/csrf/stores"
	"github.com/velocitykode/velocity/view"
)

// Configure registers the app's service providers. main.go passes this
// to v.Providers(...) — the framework calls Register on every provider
// during bootstrap, then Boot once Register has finished for all of them.
func Configure(reg *velocity.ProviderRegistry) {
	reg.Add(&AppProvider{})
}

// AppProvider wires CSRF, auth guards, and the view engine for this
// application. CSRF is built in Register so it's available to the view
// engine's shared-props closure that's set up in Boot.
type AppProvider struct{}

// Register binds the CSRF instance — runs before any provider's Boot.
func (p *AppProvider) Register(s *velocity.Services) error {
	sessionName := envOrDefault("SESSION_NAME", "velocity_session")

	csrfConfig := csrf.DefaultConfig()
	csrfConfig.Store = stores.NewSessionStore()
	csrfConfig.SessionCookieName = sessionName
	csrfConfig.ExcludePaths = []string{"/api/webhooks/*", "/health"}

	s.CSRF = csrf.New(csrfConfig)
	return nil
}

// Boot wires auth and view — runs after every provider's Register, so
// services bound by other providers are available here.
func (p *AppProvider) Boot(s *velocity.Services) error {
	if err := bootstrapAuth(s); err != nil {
		return err
	}
	return bootstrapView(s)
}

func (p *AppProvider) Shutdown(_ context.Context) error {
	return nil
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envDurationOrDefault(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}

func bootstrapAuth(s *velocity.Services) error {
	authManager := s.Auth.(*auth.Manager)
	sessionConfig := auth.NewSessionConfigFromEnv()
	provider := auth.NewORMUserProvider(s.DB.DB(), config.GetAuthModel(), authManager.GetHasher())
	sessionGuard, err := guards.NewSessionGuard(provider, sessionConfig, s.Crypto)
	if err != nil {
		return err
	}

	authManager.RegisterGuard(config.GetAuthGuard(), sessionGuard)
	return nil
}

func bootstrapView(s *velocity.Services) error {
	template, err := view.LoadTemplateFromFile(config.GetViewTemplate())
	if err != nil {
		return err
	}

	viewConfig := view.Config{
		RootTemplate: template,
		Version:      config.GetViewVersion(),
		SSREnabled:   os.Getenv("INERTIA_SSR_ENABLED") == "true",
		SSRURL:       envOrDefault("INERTIA_SSR_URL", "http://127.0.0.1:13714"),
		SSRTimeout:   envDurationOrDefault("INERTIA_SSR_TIMEOUT", 3*time.Second),
	}
	if except := os.Getenv("INERTIA_SSR_EXCEPT"); except != "" {
		for _, p := range strings.Split(except, ",") {
			if p = strings.TrimSpace(p); p != "" {
				viewConfig.SSRExcept = append(viewConfig.SSRExcept, p)
			}
		}
	}

	engine, err := view.NewEngine(viewConfig)
	if err != nil {
		return err
	}

	s.View = engine

	sessionName := envOrDefault("SESSION_NAME", "velocity_session")
	csrfInstance := s.CSRF.(*csrf.CSRF)

	engine.SetSharePropsFunc(func(r *http.Request) (view.Props, error) {
		props := view.Props{}
		if cookie, err := r.Cookie(sessionName); err == nil {
			if token, err := csrfInstance.GetToken(cookie.Value); err == nil && token != "" {
				props["csrf_token"] = token
			}
		}
		return props, nil
	})

	return nil
}
