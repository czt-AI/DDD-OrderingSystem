package service

import (
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Domain/Model"
	"DDD-OrderingSystem/Domain/Query"
)

// OrderService 订单服务接口
type OrderService interface {
	CreateOrder(ctx context.Context, customerId uint64, items []Model.OrderItem) error
	GetOrderDetails(ctx context.Context, orderID uint64) (*Model.Order, error)
}

// OrderServiceImpl 订单服务实现
type OrderServiceImpl struct {
	commandHandler Command.Handler
	queryHandler  Query.Handler
	repository    Repository.OrderRepository
}

// NewOrderServiceImpl 创建订单服务实例
func NewOrderServiceImpl(commandHandler Command.Handler, queryHandler Query.Handler, repository Repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		commandHandler: commandHandler,
		queryHandler:  queryHandler,
		repository:    repository,
	}
}

// CreateOrder 创建订单
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, customerId uint64, items []Model.OrderItem) error {
	// 创建命令
	createOrderCommand := Command.NewOrderCreateCommand(customerId, items)

	// 处理命令
	if err := s.commandHandler.Handle(ctx, createOrderCommand); err != nil {
		return err
	}

	// 获取订单ID
	orderID := createOrderCommand.GetOrderID()

	// 查询订单详情
	order, err := s.queryHandler.Handle(ctx, Query.NewOrderDetailsQuery(orderID))
	if err != nil {
		return err
	}

	return nil
}

// GetOrderDetails 获取订单详情
func (s *OrderServiceImpl) GetOrderDetails(ctx context.Context, orderID uint64) (*Model.Order, error) {
	return s.queryHandler.Handle(ctx, Query.NewOrderDetailsQuery(orderID))
}
