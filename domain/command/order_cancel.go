package command

import (
	"DDD-OrderingSystem/Domain/Model"
)

// OrderCancelCommand 取消订单命令
type OrderCancelCommand struct {
	OrderID uint64
}

// NewOrderCancelCommand 创建取消订单命令
func NewOrderCancelCommand(orderID uint64) *OrderCancelCommand {
	return &OrderCancelCommand{
		OrderID: orderID,
	}
}

// Execute 执行命令
func (c *OrderCancelCommand) Execute(ctx context.Context) error {
	// 这里可以添加取消订单的逻辑，例如更新订单状态等

	// 创建取消订单事件
	canceledEvent := Model.NewOrderCancelledEvent(c.OrderID)

	// 应用事件
	// 这里应该将事件保存到事件存储中，或者通过其他方式发布

	return nil
}
