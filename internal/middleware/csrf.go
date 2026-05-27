package middleware

import (
	"github.com/velocitykode/velocity/bond"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/router"
)

// CSRFTokenMiddleware publishes the CSRF token under the "csrfToken"
// root-template variable so app.go.html can stamp it into the
// <meta name="csrf-token"> tag. Reads through csrf.TokenForRequest so
// the value is byte-identical to any other reader on the same request
// (e.g. sharePropsFunc populating csrf_token). Must run AFTER the
// framework CSRF middleware, which attaches the request-scoped token
// cache.
//
// Publishes via bond.WithTemplateData, the only API bond's renderer
// reads when building the root-template Execute map. gonertia's own
// SetTemplateDatum lives in a different context key and never reaches
// the bond template.
func CSRFTokenMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(c *router.Context) error {
		if token, err := csrf.TokenForRequest(c.Request); err == nil && token != "" {
			ctx := bond.WithTemplateData(c.Request.Context(), "csrfToken", token)
			c.Request = c.Request.WithContext(ctx)
		}
		return next(c)
	}
}
