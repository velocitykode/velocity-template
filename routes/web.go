package routes

import (
	"{{MODULE_NAME}}/internal/handlers"
	"{{MODULE_NAME}}/internal/middleware"

	"github.com/velocitykode/velocity"
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

	r.Web(func(web router.Router) {
		// Root - / always redirects to /login.
		web.Get("/", func(c *router.Context) error {
			return c.Redirect(router.StatusFound, "/login")
		})

		// Guest-only - middleware.Guest redirects authenticated users
		// to /dashboard.
		web.Group("", func(guest router.Router) {
			guest.Get("/login", handlers.AuthShowLoginForm)        // /login
			guest.Post("/login", handlers.AuthLogin)               // /login
			guest.Get("/register", handlers.AuthShowRegisterForm)  // /register
			guest.Post("/register", handlers.AuthRegister)         // /register
		}).Use(middleware.Guest)

		web.Post("/logout", handlers.AuthLogout) // /logout

		// Authenticated - middleware.Auth redirects guests to /login.
		web.Group("", func(auth router.Router) {
			auth.Get("/dashboard", handlers.Dashboard) // /dashboard
		}).Use(middleware.Auth)
	})
}
