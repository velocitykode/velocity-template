package routes

import (
	"{{MODULE_NAME}}/app/http/controllers"
	"{{MODULE_NAME}}/app/http/middleware"

	"github.com/velocitykode/velocity/pkg/router"
)

func init() {
	router.Register(func(r router.Router) {
		// Guest routes (only accessible when NOT authenticated)
		r.Get("/login", middleware.GuestMiddleware(controllers.AuthShowLoginForm))
		r.Post("/login", middleware.GuestMiddleware(controllers.AuthLogin))
		r.Get("/register", middleware.GuestMiddleware(controllers.AuthShowRegisterForm))
		r.Post("/register", middleware.GuestMiddleware(controllers.AuthRegister))

		// Logout route (accessible to authenticated users)
		r.Post("/logout", controllers.AuthLogout)

		// Protected routes (require authentication)
		r.Get("/", middleware.AuthMiddleware(controllers.Dashboard))
		r.Get("/dashboard", middleware.AuthMiddleware(controllers.Dashboard))
	})
}
