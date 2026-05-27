package app

import (
	"{{MODULE_NAME}}/internal/middleware"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/router"
	"github.com/velocitykode/velocity/view"
)

// Middleware configures the application's middleware stacks.
//
// The framework calls this once during bootstrap with a *MiddlewareStack
// that splits middleware into three scopes:
//
//   - Global: runs on every request
//   - Web:    runs on routes inside r.Web(...)
//   - API:    runs on routes inside r.API(prefix, ...)
//
// CSRF lives on Services (framework-built during velocity.New()); the
// view engine is wired by AppProvider.Boot. The save-at-end session
// middleware is auto-installed by velocity.bootstrap (since e7a32a1)
// when the session guard is active, so the template no longer ships
// its own session wrapper.
func Middleware(m *velocity.MiddlewareStack) {
	m.Global(
		middleware.LoggingMiddleware,      // Log all requests (no framework export yet)
		middleware.TrustProxiesMiddleware, // Handle X-Forwarded-* headers (no framework export yet)
		router.CORS(router.CORSConfig{ // Framework CORS (velocity/router/cors.go)
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "X-Inertia", "X-Inertia-Version"},
			AllowCredentials: true,
		}),
		velocity.PreventRequestsDuringMaintenance(),    // Framework maintenance gate (velocity/maintenance.go)
		router.BodyLimit(10<<20),                       // Framework body-limit (velocity/router/body_limit.go) - 10MB
		middleware.TrimStringsMiddleware,               // Trim whitespace from string inputs
		middleware.ConvertEmptyStringsToNullMiddleware, // Convert "" to nil
	)

	s := m.Services()
	csrfInstance := s.CSRF.(*csrf.CSRF)
	viewEngine := s.View.(*view.Engine)

	m.Web(
		csrfInstance.RouterMiddleware(), // Attaches request-scoped token cache + validates unsafe methods
		middleware.CSRFTokenMiddleware,  // Inject CSRF token into template data (reads TokenForRequest)
		viewEngine.Middleware(),         // Inertia version + X-Inertia headers
	)

	m.API(
		middleware.EnsureJSONMiddleware, // Force JSON response content-type (sets response header; not the same as router.ContentTypeJSON which validates request headers)
	)
}
