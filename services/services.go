package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/asadbekGo/book-shop-api-gateway/config"
	pb "github.com/asadbekGo/book-shop-api-gateway/genproto/order_service"
)

type IServiceManager interface {
	OrderService() pb.OrderServiceClient
}

type serviceManager struct {
	orderService pb.OrderServiceClient
}

func (s *serviceManager) OrderService() pb.OrderServiceClient {
	return s.orderService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.OrderServiceHost, conf.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		orderService: pb.NewOrderServiceClient(connOrder),
	}

	return serviceManager, nil
}
