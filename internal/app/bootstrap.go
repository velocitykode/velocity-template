package app

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"{{MODULE_NAME}}/config"
	"{{MODULE_NAME}}/internal/models"

	"github.com/velocitykode/velocity"
	"github.com/velocitykode/velocity/bond/vite"
	"github.com/velocitykode/velocity/csrf"
	"github.com/velocitykode/velocity/view"
)

// Configure registers the app's modules. main.go passes this
// to v.Modules(...) - the framework calls Init on every module
// during bootstrap, then Start once Init has finished for all of them.
func Configure(reg *velocity.ModuleRegistry) {
	reg.Add(&AppModule{})
}

// AppModule wires this application's auth model and view engine. CSRF and
// the schemes themselves are built by the framework during velocity.New();
// what the framework cannot know is which model authenticates, so Init
// installs that and Start stands up the Inertia view engine.
type AppModule struct{}

// Init runs before any module's Start. It declares this application's
// auth model: velocity.New has already built the schemes against the
// framework's built-in user model, and SetAuthModel re-points every one of
// them at ours.
//
// This one line is where authentication chooses its model. To authenticate
// a different one - an Admin, say - change the type parameter. A model whose
// columns differ from email/password/remember_token names them:
//
//	velocity.SetAuthModel[models.Admin](s,
//	    velocity.WithAuthIdentifierColumn("username"),
//	    velocity.WithAuthPasswordColumn("pass_hash"),
//	)
func (m *AppModule) Init(s *velocity.Services) error {
	return velocity.SetAuthModel[models.User](s)
}

// Start wires the view engine - runs after every module's Init.
func (m *AppModule) Start(s *velocity.Services) error {
	return bootstrapView(s)
}

func (m *AppModule) Shutdown(_ context.Context) error {
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
		SSREnabled:   os.Getenv("VIEW_SSR_ENABLED") == "true",
		SSRURL:       envOrDefault("VIEW_SSR_URL", "http://127.0.0.1:13714"),
		SSRTimeout:   envDurationOrDefault("VIEW_SSR_TIMEOUT", 3*time.Second),
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
	if except := os.Getenv("VIEW_SSR_EXCEPT"); except != "" {
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

	engine.SetSharePropsFunc(func(r *http.Request) (view.Props, error) {
		props := view.Props{}
		if token, err := csrf.TokenForRequest(r); err == nil && token != "" {
			props["csrf_token"] = token
		}
		return props, nil
	})

	return nil
}
