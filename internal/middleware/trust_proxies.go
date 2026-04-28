package middleware

import (
	"strings"

	"github.com/velocitykode/velocity/router"
)

// TrustProxiesMiddleware handles X-Forwarded-* headers from trusted proxies
func TrustProxiesMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		// Handle X-Forwarded-For header
		if xff := c.Request.Header.Get("X-Forwarded-For"); xff != "" {
			// Get first IP from comma-separated list
			ips := strings.Split(xff, ",")
			if len(ips) > 0 {
				c.Request.RemoteAddr = strings.TrimSpace(ips[0])
			}
		}

		// Handle X-Forwarded-Host header
		if host := c.Request.Header.Get("X-Forwarded-Host"); host != "" {
			c.Request.Host = host
		}

		// X-Forwarded-Proto is intentionally NOT applied to r.URL.Scheme.
		// On a server-received request r.URL is path-only, so r.URL.String()
		// returns just the path - which is what Inertia's "url" page prop
		// expects. Setting Scheme without Host turns r.URL.String() into
		// "https:///path" (empty host) and breaks Inertia navigation.
		// Code that genuinely needs scheme awareness should read the
		// X-Forwarded-Proto header directly.

		return next(c)
	}
}
