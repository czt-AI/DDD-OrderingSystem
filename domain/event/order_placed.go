package event

import (
	"DDD-OrderingSystem/Domain/Model"
	"time"
)

// OrderPlacedEvent 订单已下单事件
type OrderPlacedEvent struct {
	OrderID    uint64
	CustomerID uint64
	Items      []Model.OrderItem
	Status     Model.OrderStatus
	CreatedAt  time.Time
}

// NewOrderPlacedEvent 创建订单已下单事件
func NewOrderPlacedEvent(orderID uint64, customerID uint64, items []Model.OrderItem, status Model.OrderStatus) *OrderPlacedEvent {
	return &OrderPlacedEvent{
		OrderID:    orderID,
		CustomerID: customerID,
		Items:      items,
		Status:     status,
		CreatedAt:  time.Now(),
	}
}
