package config

// GetViewTemplate returns the view template path (read at call time)
func GetViewTemplate() string {
	return envOr("VIEW_TEMPLATE", "resources/views/app.go.html")
}

// GetViewVersion returns the view version (read at call time)
func GetViewVersion() string {
	return envOr("VIEW_VERSION", "1.0")
}
