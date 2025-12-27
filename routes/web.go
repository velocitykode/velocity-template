package routes

import (
	"{{MODULE_NAME}}/app/http/controllers"
	"{{MODULE_NAME}}/app/http/middleware"

	"github.com/velocitykode/velocity/pkg/router"
)

func init() {
	router.Register(func(r router.Router) {
		// Create controller instances
		authController := controllers.NewAuthController()
		dashboardController := controllers.NewDashboardController()

		// Guest routes (only accessible when NOT authenticated)
		r.Get("/login", middleware.GuestMiddleware(authController.ShowLoginForm))
		r.Post("/login", middleware.GuestMiddleware(authController.Login))
		r.Get("/register", middleware.GuestMiddleware(authController.ShowRegisterForm))
		r.Post("/register", middleware.GuestMiddleware(authController.Register))

		// Logout route (accessible to authenticated users)
		r.Post("/logout", authController.Logout)

		// Protected routes (require authentication)
		r.Get("/", middleware.AuthMiddleware(dashboardController.Index))
		r.Get("/dashboard", middleware.AuthMiddleware(dashboardController.Index))
	})
}
