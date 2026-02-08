package handlers

import (
	"net/http"

	"github.com/velocitykode/velocity/pkg/router"
)

func Health(ctx *router.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"status": "healthy",
	})
}
