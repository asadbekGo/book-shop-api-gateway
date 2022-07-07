package main

import (
	"fmt"

	"github.com/asadbekGo/book-shop-api-gateway/api"
	"github.com/asadbekGo/book-shop-api-gateway/config"
	"github.com/asadbekGo/book-shop-api-gateway/pkg/logger"
	"github.com/asadbekGo/book-shop-api-gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	fmt.Println("Listening...")

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
