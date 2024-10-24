package grpc

import (
	"context"
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Domain/Model"
	"DDD-OrderingSystem/Infrastructure/Repository"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"orderingservice"
)

// OrderServiceServer gRPC订单服务服务器
type OrderServiceServer struct {
	repository Repository.OrderRepository
}

// NewOrderServiceServer 创建新的OrderServiceServer实例
func NewOrderServiceServer(repository Repository.OrderRepository) *OrderServiceServer {
	return &OrderServiceServer{
		repository: repository,
	}
}

// CreateOrder 创建订单
func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *orderingservice.CreateOrderRequest) (*orderingservice.CreateOrderResponse, error) {
	// 将请求转换为领域模型
	order := &Model.Order{
		CustomerID: req.CustomerId,
		Items:      make([]Model.OrderItem, len(req.Items)),
		Status:     Model.OrderStatus(req.Status),
	}
	for i, item := range req.Items {
		order.Items[i] = Model.OrderItem{
			ProductID: item.ProductId,
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
		}
	}

	// 创建命令
	createOrderCommand := Command.NewOrderCreateCommand(order.CustomerID, order.Items)

	// 处理命令
	if err := Command.Handle(ctx, createOrderCommand); err != nil {
		return nil, err
	}

	// 返回响应
	return &orderingservice.CreateOrderResponse{
		OrderId: createOrderCommand.GetOrderID(),
	}, nil
}

// GetOrder 获取订单
func (s *OrderServiceServer) GetOrder(ctx context.Context, req *orderingservice.GetOrderRequest) (*orderingservice.GetOrderResponse, error) {
	order, err := s.repository.FindById(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	// 将领域模型转换为gRPC响应
	response := &orderingservice.GetOrderResponse{
		OrderId: order.ID,
		CustomerId: order.CustomerID,
		Items:     make([]*orderingservice.OrderItem, len(order.Items)),
		Status:    order.Status,
	}
	for i, item := range order.Items {
		response.Items[i] = &orderingservice.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
		}
	}

	return response, nil
}
