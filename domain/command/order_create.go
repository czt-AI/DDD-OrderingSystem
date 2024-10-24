package command

import (
	"DDD-OrderingSystem/Domain/Model"
)

// OrderCreateCommand 创建订单命令
type OrderCreateCommand struct {
	CustomerID uint64
	Items      []Model.OrderItem
}

// NewOrderCreateCommand 创建订单创建命令
func NewOrderCreateCommand(customerID uint64, items []Model.OrderItem) *OrderCreateCommand {
	return &OrderCreateCommand{
		CustomerID: customerID,
		Items:      items,
	}
}

// Execute 执行命令
func (c *OrderCreateCommand) Execute(ctx context.Context) error {
	// 这里可以添加创建订单的逻辑，例如生成订单ID，设置订单状态等

	// 创建订单模型
	order := Model.Order{
		CustomerID: c.CustomerID,
		Items:      c.Items,
		Status:     Model.OrderStatusNew,
	}

	// 应用事件
事件的创建逻辑应该在这里实现，例如：
	placedEvent := Model.NewOrderPlacedEvent(order.ID, order.CustomerID, order.Items, order.Status)
	// 这里应该将事件保存到事件存储中，或者通过其他方式发布

	return nil
}

// GetOrderID 获取订单ID
func (c *OrderCreateCommand) GetOrderID() uint64 {
	// 这里应该返回订单ID，实际应用中这个值应该由执行命令的逻辑来设置
	return 0
}
