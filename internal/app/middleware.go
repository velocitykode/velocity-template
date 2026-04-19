package app

import (
	"{{MODULE_NAME}}/internal/middleware"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/csrf"
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
// view engine is wired by AppProvider.Boot.
func Middleware(m *velocity.MiddlewareStack) {
	m.Global(
		middleware.LoggingMiddleware,                          // Log all requests
		middleware.TrustProxiesMiddleware,                     // Handle X-Forwarded-* headers
		middleware.CORSMiddleware,                             // CORS preflight + headers
		middleware.PreventRequestsDuringMaintenanceMiddleware, // 503 when in maintenance mode
		middleware.ValidatePostSizeMiddleware(10<<20),         // Reject requests > 10MB
		middleware.TrimStringsMiddleware,                      // Trim whitespace from string inputs
		middleware.ConvertEmptyStringsToNullMiddleware,        // Convert "" to nil
	)

	s := m.Services()
	csrfInstance := s.CSRF.(*csrf.CSRF)
	viewEngine := s.View.(*view.Engine)

	m.Web(
		middleware.SessionMiddleware,    // Session cookie (must run before CSRF)
		middleware.CSRFTokenMiddleware,  // Inject CSRF token into template data
		csrfInstance.RouterMiddleware(), // Validate CSRF token on unsafe methods
		viewEngine.Middleware(),         // Inertia version + X-Inertia headers
	)

	m.API(
		middleware.EnsureJSONMiddleware, // Force JSON content-type on responses
	)
}
