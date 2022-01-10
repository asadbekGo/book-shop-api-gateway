package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/asadbekGo/book-shop-api-gateway/config"
	pbCatalog "github.com/asadbekGo/book-shop-api-gateway/genproto/catalog_service"
	pbOrder "github.com/asadbekGo/book-shop-api-gateway/genproto/order_service"
)

type IServiceManager interface {
	OrderService() pbOrder.OrderServiceClient
	CatalogService() pbCatalog.CatalogServiceClient
}

type serviceManager struct {
	orderService   pbOrder.OrderServiceClient
	catalogService pbCatalog.CatalogServiceClient
}

func (s *serviceManager) OrderService() pbOrder.OrderServiceClient {
	return s.orderService
}

func (s *serviceManager) CatalogService() pbCatalog.CatalogServiceClient {
	return s.catalogService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.OrderServiceHost, conf.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CatalogServiceHost, conf.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		orderService:   pbOrder.NewOrderServiceClient(connOrder),
		catalogService: pbCatalog.NewCatalogServiceClient(connCatalog),
	}

	return serviceManager, nil
}
