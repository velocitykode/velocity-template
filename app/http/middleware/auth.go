package middleware

import (
	"net/http"

	"github.com/velocitykode/velocity/pkg/auth"
	"github.com/velocitykode/velocity/pkg/router"
)

// AuthMiddleware redirects to login if not authenticated
func AuthMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) error {
		if !auth.Check(ctx.Request) {
			return ctx.Redirect(http.StatusSeeOther, "/login")
		}
		return next(ctx)
	}
}

// GuestMiddleware redirects to dashboard if already authenticated
func GuestMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx *router.Context) error {
		if auth.Check(ctx.Request) {
			return ctx.Redirect(http.StatusSeeOther, "/dashboard")
		}
		return next(ctx)
	}
}
