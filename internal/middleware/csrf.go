package middleware

import (
	"github.com/romsar/gonertia"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/router"
)

// CSRFTokenMiddleware sets the CSRF token in the template data for the meta tag.
// Reads via csrf.TokenForRequest so the value is byte-identical to any other
// reader on the same request (e.g. sharePropsFunc populating csrf_token).
// Must run AFTER the framework CSRF middleware, which attaches the
// request-scoped token cache.
func CSRFTokenMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		ctx := c.Request.Context()
		if token, err := csrf.TokenForRequest(c.Request); err == nil && token != "" {
			ctx = gonertia.SetTemplateDatum(ctx, "csrfToken", token)
			c.Request = c.Request.WithContext(ctx)
		}
		return next(c)
	}
}
