package config

// GetAuthGuard returns the auth guard (read at call time)
func GetAuthGuard() string {
	return envOr("AUTH_GUARD", "web")
}

// GetAuthModel returns the auth model (read at call time)
func GetAuthModel() string {
	return envOr("AUTH_MODEL", "User")
}
