package controllers

import (
	"github.com/velocitykode/velocity/pkg/auth"
	"github.com/velocitykode/velocity/pkg/router"
	"github.com/velocitykode/velocity/pkg/view"
)

type DashboardController struct{}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

// Index displays the dashboard
func (c *DashboardController) Index(ctx *router.Context) error {
	user := auth.User(ctx.Request)

	// Convert user to map for props
	userMap := make(map[string]interface{})
	if mockUser, ok := user.(*auth.MockUser); ok {
		userMap["id"] = mockUser.ID
		userMap["name"] = mockUser.Name
		userMap["email"] = mockUser.Email
	}

	view.Render(ctx.Response, ctx.Request, "Dashboard", view.Props{
		"auth": map[string]interface{}{
			"user": userMap,
		},
	})
	return nil
}
