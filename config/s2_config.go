package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type ConfigOS struct {
	Environment string // develop, staging, production

	OrderServiceHost string
	OrderServicePort int

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func LoadOS() ConfigOS {
	c := ConfigOS{}

	c.Environment = cast.ToString(getOrReturnDefaultOS("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefaultOS("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefaultOS("HTTP_PORT", ":8080"))
	c.OrderServiceHost = cast.ToString(getOrReturnDefaultOS("ORDER_SERVICE_HOST", "127.0.0.1"))
	c.OrderServicePort = cast.ToInt(getOrReturnDefaultOS("ORDER_SERVICE_PORT", 50052))

	c.CtxTimeout = cast.ToInt(getOrReturnDefaultOS("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefaultOS(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
