package api

import (
	"github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"

	//_ "github.com/asadbekGo/book-shop-api-gateway/api/docs" // swag
	v1 "github.com/asadbekGo/book-shop-api-gateway/api/handlers/v1"
	"github.com/asadbekGo/book-shop-api-gateway/config"
	"github.com/asadbekGo/book-shop-api-gateway/pkg/logger"
	"github.com/asadbekGo/book-shop-api-gateway/services"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/os_v1")
	api.POST("/orders", handlerV1.CreateOrder)
	api.GET("/orders/:id", handlerV1.GetOrder)
	api.GET("/orders", handlerV1.ListOrders)
	api.PUT("/orders/:id", handlerV1.UpdateOrder)
	api.DELETE("/orders/:id", handlerV1.DeleteOrder)

	// url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}