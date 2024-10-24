package grpc

import (
	"context"
	"DDD-OrderingSystem/OrderingApplication/Domain/Model"
	"DDD-OrderingSystem/OrderingApplication/Infrastructure/Adapters/HealthCheck"
	"google.golang.org/grpc"
	"orderingservice"
	"userservice"
	"healthcheckservice"
)

// GRPCClient gRPC客户端
type GRPCClient struct {
	OrderClient    orderingservice.OrderServiceClient
	UserClient     userservice.UserServiceClient
	HealthClient   healthcheckservice.HealthCheckServiceClient
}

// NewGRPCClient 创建新的gRPC客户端实例
func NewGRPCClient(host string) (*GRPCClient, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &GRPCClient{
		OrderClient:    orderingservice.NewOrderServiceClient(conn),
		UserClient:     userservice.NewUserServiceClient(conn),
		HealthClient:   healthcheckservice.NewHealthCheckServiceClient(conn),
	}, nil
}

// CreateOrder 创建订单
func (c *GRPCClient) CreateOrder(ctx context.Context, customerId uint64, items []Model.OrderItem) (*orderingservice.CreateOrderResponse, error) {
	req := &orderingservice.CreateOrderRequest{
		CustomerId: customerId,
		Items:      make([]*orderingservice.OrderItem, len(items)),
	}
	for i, item := range items {
		req.Items[i] = &orderingservice.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
		}
	}
	return c.OrderClient.CreateOrder(ctx, req)
}

// GetUser 获取用户
func (c *GRPCClient) GetUser(ctx context.Context, userId uint64) (*userservice.GetUserResponse, error) {
	return c.UserClient.GetUser(ctx, &userservice.GetUserRequest{UserId: userId})
}

// CheckHealth 检查服务健康状态
func (c *GRPCClient) CheckHealth(ctx context.Context) (*healthcheckservice.CheckResponse, error) {
	return c.HealthClient.Check(ctx, &healthcheckservice.CheckRequest{})
}