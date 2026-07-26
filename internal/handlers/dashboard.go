package handlers

import (
	"{{MODULE_NAME}}/internal/models"

	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/router"
	"github.com/velocitykode/velocity/view"
)

// Dashboard displays the dashboard
func Dashboard(ctx *router.Context) error {
	user := auth.FromContext(ctx).User(ctx.Request)

	// Convert user to map for props. The concrete type is *models.User
	// because internal/app/bootstrap.go installs the auth provider on that
	// model - change the model there and change this assertion with it.
	userMap := make(map[string]interface{})
	if authUser, ok := user.(*models.User); ok {
		userMap["id"] = authUser.ID
		userMap["name"] = authUser.Name
		userMap["email"] = authUser.Email
	}

	view.Render(ctx, "Dashboard", view.Props{
		"auth": map[string]interface{}{
			"user": userMap,
		},
	})
	return nil
}
