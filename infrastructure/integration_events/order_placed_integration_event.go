package integrationevents

import (
	"DDD-OrderingSystem/Domain/Model"
	"time"
)

// OrderPlacedIntegrationEvent 订单已下单集成事件
type OrderPlacedIntegrationEvent struct {
	OrderID    uint64
	CustomerID uint64
	Items      []OrderItem
	Status     Model.OrderStatus
	CreatedAt  time.Time
}

// OrderItem 订单项结构
type OrderItem struct {
	ProductID uint64
	Quantity  int
	Price     float64
}

// NewOrderPlacedIntegrationEvent 创建订单已下单集成事件
func NewOrderPlacedIntegrationEvent(orderID uint64, customerID uint64, items []OrderItem, status Model.OrderStatus) *OrderPlacedIntegrationEvent {
	return &OrderPlacedIntegrationEvent{
		OrderID:    orderID,
		CustomerID: customerID,
		Items:      items,
		Status:     status,
		CreatedAt:  time.Now(),
	}
}
