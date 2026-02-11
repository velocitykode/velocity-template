package routes

import (
	"{{MODULE_NAME}}/internal/handlers"
	"{{MODULE_NAME}}/internal/middleware"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/pkg/router"
)

func Register(v *velocity.App) {
	r := v.Router

	r.Get("/health", handlers.Health)

	r.Group("", func(guest router.Router) {
		guest.Get("/login", handlers.AuthShowLoginForm)
		guest.Post("/login", handlers.AuthLogin)
		guest.Get("/register", handlers.AuthShowRegisterForm)
		guest.Post("/register", handlers.AuthRegister)
	}).Use(middleware.Guest)

	r.Post("/logout", handlers.AuthLogout)

	r.Group("", func(auth router.Router) {
		auth.Get("/", handlers.Dashboard)
		auth.Get("/dashboard", handlers.Dashboard)
	}).Use(middleware.Auth)
}
