package app

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"{{MODULE_NAME}}/config"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/bond/vite"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/view"
)

// Configure registers the app's service providers. main.go passes this
// to v.Providers(...) - the framework calls Register on every provider
// during bootstrap, then Boot once Register has finished for all of them.
func Configure(reg *velocity.ProviderRegistry) {
	reg.Add(&AppProvider{})
}

// AppProvider wires the view engine for this application. CSRF and auth
// guards are built by the framework from env vars during velocity.New(),
// so we only need to stand up the Inertia view engine and share the
// CSRF token into its props.
type AppProvider struct{}

// Register runs before any provider's Boot. Nothing to do here - the
// framework already built CSRF, Auth, and the session guard.
func (p *AppProvider) Register(s *velocity.Services) error {
	return nil
}

// Boot wires the view engine - runs after every provider's Register.
func (p *AppProvider) Boot(s *velocity.Services) error {
	return bootstrapView(s)
}

func (p *AppProvider) Shutdown(_ context.Context) error {
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

func bootstrapView(s *velocity.Services) error {
	// view.Config.RootTemplate takes the HTML content string, not a path.
	// Read the file ourselves so bond can parse + validate it.
	templateBytes, err := os.ReadFile(config.GetViewTemplate())
	if err != nil {
		return err
	}

	// Vite helper exposes {{ vite "resources/js/app.tsx" }} to the
	// root template. In dev (public/hot exists) it emits dev-server
	// tags; in prod it reads public/build/.vite/manifest.json and
	// emits hashed asset URLs with modulepreload + stylesheet links.
	viteHelper := vite.New()

	viewConfig := view.Config{
		RootTemplate: string(templateBytes),
		Version:      config.GetViewVersion(),
		SSREnabled:   os.Getenv("INERTIA_SSR_ENABLED") == "true",
		SSRURL:       envOrDefault("INERTIA_SSR_URL", "http://127.0.0.1:13714"),
		SSRTimeout:   envDurationOrDefault("INERTIA_SSR_TIMEOUT", 3*time.Second),
		Funcs: template.FuncMap{
			"vite": func(entrypoints ...string) template.HTML {
				out, _ := viteHelper.Tags(entrypoints...)
				return out
			},
			// React Fast Refresh preamble - emits a script in dev
			// mode, empty in prod. Must precede {{ vite ... }} so the
			// preamble runs before @vite/client.
			"viteReactRefresh": func() template.HTML {
				out, _ := viteHelper.ReactRefreshTag()
				return out
			},
		},
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

	s.View = engine

	sessionName := envOrDefault("SESSION_NAME", "velocity_session")
	csrfInstance := s.CSRF.(*csrf.CSRF)

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
