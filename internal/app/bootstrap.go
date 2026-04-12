package app

import (
	"net/http"
	"os"
	"strings"
	"time"

	"{{MODULE_NAME}}/config"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/auth"
	"github.com/velocitykode/velocity/auth/drivers/guards"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/csrf/stores"
	"github.com/velocitykode/velocity/view"
)

// Bootstrap configures app-specific services on the Velocity app instance.
// Core services (crypto, ORM, logger, cache, events, queue, storage) are
// already initialized by velocity.Default().
func Bootstrap(v *velocity.App) error {
	// 1. Register auth guards (app-specific: session guard with user model)
	if err := bootstrapAuth(v); err != nil {
		return err
	}

	// 2. Configure CSRF with session store (app-specific)
	bootstrapCSRF(v)

	// 3. Configure view engine with template and shared props (app-specific)
	if err := bootstrapView(v); err != nil {
		return err
	}

	// 4. Register event listeners (app-specific)
	initEvents(v)

	// 5. Apply middleware to the router
	bootstrapMiddleware(v)

	// 6. Serve static files
	v.Router.Static("public")

	return nil
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envDurationOrDefault(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}

func bootstrapAuth(v *velocity.App) error {
	authManager := v.Auth.(*auth.Manager)
	sessionConfig := auth.NewSessionConfigFromEnv()
	provider := auth.NewORMUserProvider(v.DB.DB(), config.GetAuthModel(), authManager.GetHasher())
	sessionGuard, err := guards.NewSessionGuard(provider, sessionConfig, v.Crypto)
	if err != nil {
		return err
	}

	authManager.RegisterGuard(config.GetAuthGuard(), sessionGuard)
	return nil
}

func bootstrapCSRF(v *velocity.App) {
	sessionName := os.Getenv("SESSION_NAME")
	if sessionName == "" {
		sessionName = "velocity_session"
	}

	csrfConfig := csrf.DefaultConfig()
	csrfConfig.Store = stores.NewSessionStore()
	csrfConfig.SessionCookieName = sessionName
	csrfConfig.ExcludePaths = []string{"/api/webhooks/*", "/health"}

	v.CSRF = csrf.New(csrfConfig)
}

func bootstrapView(v *velocity.App) error {
	template, err := view.LoadTemplateFromFile(config.GetViewTemplate())
	if err != nil {
		return err
	}

	viewConfig := view.Config{
		RootTemplate: template,
		Version:      config.GetViewVersion(),
		SSREnabled:   os.Getenv("INERTIA_SSR_ENABLED") == "true",
		SSRURL:       envOrDefault("INERTIA_SSR_URL", "http://127.0.0.1:13714"),
		SSRTimeout:   envDurationOrDefault("INERTIA_SSR_TIMEOUT", 3*time.Second),
	}
	if except := os.Getenv("INERTIA_SSR_EXCEPT"); except != "" {
		for _, p := range strings.Split(except, ",") {
			if p = strings.TrimSpace(p); p != "" {
				viewConfig.SSRExcept = append(viewConfig.SSRExcept, p)
			}
		}
	}

	engine, err := view.NewEngine(viewConfig)
	if err != nil {
		return err
	}

	v.View = engine

	sessionName := os.Getenv("SESSION_NAME")
	if sessionName == "" {
		sessionName = "velocity_session"
	}

	csrfInstance := v.CSRF.(*csrf.CSRF)
	engine.SetSharePropsFunc(func(r *http.Request) (view.Props, error) {
		props := view.Props{}
		if cookie, err := r.Cookie(sessionName); err == nil {
			if token, err := csrfInstance.GetToken(cookie.Value); err == nil && token != "" {
				props["csrf_token"] = token
			}
		}
		return props, nil
	})

	return nil
}

func bootstrapMiddleware(v *velocity.App) {
	stacks := GetMiddlewareStacks(v)

	for _, mw := range stacks.Global {
		v.Router.Use(mw)
	}
	for _, mw := range stacks.Web {
		v.Router.Use(mw)
	}
}
