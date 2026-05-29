package handlers

import (
	"{{MODULE_NAME}}/internal/models"

	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/router"
	"github.com/velocitykode/velocity/validation"
	"github.com/velocitykode/velocity/validation/vform"
	"github.com/velocitykode/velocity/view"
)

// LoginRequest is the form-request schema for POST /login. vform.Form[T]
// binds the request body into a *LoginRequest, runs Rules(), flashes
// errors via WithErrors/WithInput, redirects back, and returns
// router.ErrValidationAborted to short-circuit the handler. The frontend
// reads the flashed errors/old props on the next render.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func (r *LoginRequest) Rules() validation.Rules {
	return validation.Rules{
		"email":    {"required", "email"},
		"password": {"required"},
	}
}

// RegisterRequest schema for POST /register. The unique:users,email rule
// is enforced by the validation engine against the configured database
// driver; the confirmed rule pairs password with password_confirmation
// without a manual equality check.
type RegisterRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (r *RegisterRequest) Rules() validation.Rules {
	return validation.Rules{
		"name":     {"required", "max:255"},
		"email":    {"required", "email", "unique:users,email"},
		"password": {"required", "min:8", "confirmed"},
	}
}

// AuthShowLoginForm displays the login page
func AuthShowLoginForm(ctx *router.Context) error {
	view.Render(ctx, "Auth/Login", view.Props{})
	return nil
}

// AuthLogin handles the login request
func AuthLogin(ctx *router.Context) error {
	req, err := vform.Form[LoginRequest](ctx)
	if err != nil {
		return err
	}

	credentials := map[string]interface{}{
		"email":    req.Email,
		"password": req.Password,
	}

	success, _ := auth.FromContext(ctx).Attempt(ctx.Response, ctx.Request, credentials, req.Remember)
	if !success {
		ctx.WithErrors(map[string][]string{
			"email": {"These credentials do not match our records."},
		})
		ctx.WithInput(map[string]any{"email": req.Email})
		view.Back(ctx)
		return nil
	}

	// Honour the intended destination the auth middleware stashed in the
	// session before bouncing the guest to /login (falls back to
	// /dashboard for a direct login).
	view.Redirect(ctx, ctx.Intended("/dashboard"))
	return nil
}

// AuthLogout handles the logout request
func AuthLogout(ctx *router.Context) error {
	auth.FromContext(ctx).Logout(ctx.Response, ctx.Request)
	view.Redirect(ctx, "/login")
	return nil
}

// AuthShowRegisterForm displays the registration page
func AuthShowRegisterForm(ctx *router.Context) error {
	view.Render(ctx, "Auth/Register", view.Props{})
	return nil
}

// AuthRegister handles the registration request
func AuthRegister(ctx *router.Context) error {
	req, err := vform.Form[RegisterRequest](ctx)
	if err != nil {
		return err
	}

	hashedPassword, err := auth.FromContext(ctx).Hash(req.Password)
	if err != nil {
		ctx.Log().Error("Failed to hash password", "error", err)
		ctx.WithErrors(map[string][]string{"password": {"Failed to process password."}})
		ctx.WithInput(map[string]any{"name": req.Name, "email": req.Email})
		view.Back(ctx)
		return nil
	}

	user, err := models.User{}.Create(ctx.Request.Context(), map[string]any{
		"name":     req.Name,
		"email":    req.Email,
		"password": hashedPassword,
	})
	if err != nil {
		ctx.Log().Error("Failed to create user", "error", err)
		ctx.WithErrors(map[string][]string{"email": {"Failed to create account. Please try again."}})
		ctx.WithInput(map[string]any{"name": req.Name, "email": req.Email})
		view.Back(ctx)
		return nil
	}

	ctx.Log().Info("User created successfully", "email", user.Email, "id", user.ID)

	credentials := map[string]interface{}{
		"email":    req.Email,
		"password": req.Password,
	}
	if success, _ := auth.FromContext(ctx).Attempt(ctx.Response, ctx.Request, credentials, false); success {
		view.Redirect(ctx, ctx.Intended("/dashboard"))
	} else {
		view.Redirect(ctx, "/login")
	}
	return nil
}
