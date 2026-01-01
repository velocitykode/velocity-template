package routes

import (
	"{{MODULE_NAME}}/app/http/controllers"
	"{{MODULE_NAME}}/app/http/middleware"

	"github.com/velocitykode/velocity/pkg/router"
)

func init() {
	router.Register(func(r router.Router) {
		// Guest routes (only accessible when NOT authenticated)
		r.Group("", func(guest router.Router) {
			guest.Use(middleware.Guest)
			guest.Get("/login", controllers.AuthShowLoginForm)
			guest.Post("/login", controllers.AuthLogin)
			guest.Get("/register", controllers.AuthShowRegisterForm)
			guest.Post("/register", controllers.AuthRegister)
		})

		// Public routes
		r.Post("/logout", controllers.AuthLogout)

		// Protected routes (require authentication)
		r.Group("", func(auth router.Router) {
			auth.Use(middleware.Auth)
			auth.Get("/", controllers.Dashboard)
			auth.Get("/dashboard", controllers.Dashboard)
		})
	})
}
