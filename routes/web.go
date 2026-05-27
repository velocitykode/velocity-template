package routes

import (
	"{{MODULE_NAME}}/internal/handlers"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/router"
)

// Register defines all application routes. main.go passes this function
// to v.Routes(...). The framework calls it with a *velocity.Routing
// already wired with the configured middleware stacks.
func Register(r *velocity.Routing) {
	// Operational endpoints sit at the top level - they don't run any
	// middleware stack so load balancers can probe /health cheaply.
	r.Health("/health")
	r.Router().StaticFallback("public")

	// Auth manager from Services - powers the framework's
	// auth.AuthMiddleware / auth.GuestMiddlewareWithRedirect (replaces
	// the template's hand-rolled middleware.Auth / middleware.Guest).
	authManager := r.Services().Auth.(*auth.Manager)

	r.Web(func(web router.Router) {
		// Root - / always redirects to /login.
		web.Get("/", func(c *router.Context) error {
			return c.Redirect(router.StatusFound, "/login")
		})

		// Guest-only - framework redirects authenticated HTML callers
		// to /dashboard (and returns 403 JSON to XHR callers).
		web.Group("", func(guest router.Router) {
			guest.Get("/login", handlers.AuthShowLoginForm)       // /login
			guest.Post("/login", handlers.AuthLogin)              // /login
			guest.Get("/register", handlers.AuthShowRegisterForm) // /register
			guest.Post("/register", handlers.AuthRegister)        // /register
		}).Use(auth.GuestMiddlewareWithRedirect(authManager, "/dashboard"))

		web.Post("/logout", handlers.AuthLogout) // /logout

		// Authenticated - framework redirects guests to /login with
		// an "intended" query param so post-login redirect works.
		web.Group("", func(authed router.Router) {
			authed.Get("/dashboard", handlers.Dashboard) // /dashboard
		}).Use(auth.AuthMiddleware(authManager))
	})
}
