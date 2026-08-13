package config

// GetAuthScheme returns the auth scheme (read at call time)
func GetAuthScheme() string {
	return envOr("AUTH_SCHEME", "web")
}
