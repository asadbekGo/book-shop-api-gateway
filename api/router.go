package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/asadbekGo/book-shop-api-gateway/api/docs" // swag
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

	api := router.Group("/v1")
	api.POST("/author", handlerV1.CreateAuthor)
	api.GET("/authors/:id", handlerV1.GetAuthor)
	api.GET("/authors", handlerV1.GetAuthors)
	api.PUT("/author/:id", handlerV1.UpdateAuthor)
	api.DELETE("/author/:id", handlerV1.DeleteAuthor)

	api.POST("/category", handlerV1.CreateCategory)
	api.GET("/categories/:id", handlerV1.GetCategory)
	api.GET("/categories", handlerV1.GetCategories)
	api.PUT("/category/:id", handlerV1.UpdateCategory)
	api.DELETE("/category/:id", handlerV1.DeleteCategory)

	api.POST("/book", handlerV1.CreateBook)
	api.GET("/books/:id", handlerV1.GetBook)
	api.GET("/books", handlerV1.GetBooks)
	api.PUT("/book/:id", handlerV1.UpdateBook)
	api.DELETE("/book/:id", handlerV1.DeleteBook)

	api.POST("/order", handlerV1.CreateOrder)
	api.GET("/orders/:id", handlerV1.GetOrder)
	api.GET("/orders", handlerV1.ListOrders)
	api.PUT("/order/:id", handlerV1.UpdateOrder)
	api.DELETE("/order/:id", handlerV1.DeleteOrder)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
