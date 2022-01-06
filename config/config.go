package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production
	// catalog ...
	CatalogServiceHost string
	CatalogServicePort int

	// order ...
	OrderServiceHost string
	OrderServicePort int

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	// api-gateway idfs
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	// catalog-service idfs
	c.CatalogServiceHost = cast.ToString(getOrReturnDefault("CATALOG_SERVICE_HOST", "127.0.0.1"))
	c.CatalogServicePort = cast.ToInt(getOrReturnDefault("CATALOG_SERVICE_PORT", 5005))

	// order-service idfs
	c.OrderServiceHost = cast.ToString(getOrReturnDefault("ORDER_SERVICE_HOST", "127.0.0.1"))
	c.OrderServicePort = cast.ToInt(getOrReturnDefault("ORDER_SERVICE_PORT", 50051))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
