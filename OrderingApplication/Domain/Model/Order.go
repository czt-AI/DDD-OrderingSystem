package model

import (
	"time"
)

// Order 订单模型
type Order struct {
	ID        uint64
	CustomerID uint64
	Items     []OrderItem
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OrderItem 订单项模型
type OrderItem struct {
	ID        uint64
	ProductID uint64
	Quantity  int
	Price     float64
}

// OrderStatus 订单状态枚举
type OrderStatus string

const (
	// OrderStatusNew 新订单
	OrderStatusNew OrderStatus = "NEW"
	// OrderStatusProcessing 处理中
	OrderStatusProcessing OrderStatus = "PROCESSING"
	// OrderStatusCompleted 完成订单
	OrderStatusCompleted OrderStatus = "COMPLETED"
	// OrderStatusCancelled 取消订单
	OrderStatusCancelled OrderStatus = "CANCELLED"
)
