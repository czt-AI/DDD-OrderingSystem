package service

import (
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Domain/Model"
)

// OrderService 订单服务接口
type OrderService interface {
	CreateOrder(ctx context.Context, customerId uint64, items []Model.OrderItem) error
}

// OrderServiceImpl 订单服务实现
type OrderServiceImpl struct {
	repository Repository.OrderRepository
}

// NewOrderServiceImpl 创建订单服务实例
func NewOrderServiceImpl(repository Repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		repository: repository,
	}
}

// CreateOrder 创建订单
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, customerId uint64, items []Model.OrderItem) error {
	// 创建订单模型
	order := Model.Order{
		CustomerID: customerId,
		Items:      items,
		Status:     Model.OrderStatusNew,
	}

	// 创建命令
	createOrderCommand := Command.NewCreateOrderCommand(customerId, items)

	// 执行命令
	if err := createOrderCommand.Execute(ctx); err != nil {
		return err
	}

	// 保存订单到数据库
	if err := s.repository.Save(ctx, &order); err != nil {
		return err
	}

	return nil
}
