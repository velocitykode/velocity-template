package routes

import (
	"{{MODULE_NAME}}/internal/handlers"
	"{{MODULE_NAME}}/internal/middleware"

	"github.com/velocitykode/velocity/pkg/router"
)

func init() {
	router.Register(func(r router.Router) {
		// Guest routes (only accessible when NOT authenticated)
		r.Group("", func(guest router.Router) {
			guest.Use(middleware.Guest)
			guest.Get("/login", handlers.AuthShowLoginForm)
			guest.Post("/login", handlers.AuthLogin)
			guest.Get("/register", handlers.AuthShowRegisterForm)
			guest.Post("/register", handlers.AuthRegister)
		})

		// Public routes
		r.Post("/logout", handlers.AuthLogout)

		// Protected routes (require authentication)
		r.Group("", func(auth router.Router) {
			auth.Use(middleware.Auth)
			auth.Get("/", handlers.Dashboard)
			auth.Get("/dashboard", handlers.Dashboard)
		})
	})
}
