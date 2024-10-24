package query

import (
	"DDD-OrderingSystem/Domain/Model"
)

// OrderDetailsQuery 订单详情查询
type OrderDetailsQuery struct {
	OrderID uint64
}

// NewOrderDetailsQuery 创建订单详情查询实例
func NewOrderDetailsQuery(orderID uint64) *OrderDetailsQuery {
	return &OrderDetailsQuery{
		OrderID: orderID,
	}
}

// Execute 执行查询
func (q *OrderDetailsQuery) Execute(ctx context.Context) (*Model.Order, error) {
	// 这里应该实现查询订单详情的逻辑，通常是通过仓库来获取数据
	// 以下代码仅为示例，实际应用中需要根据具体逻辑来获取订单详情

	order, err := Model.FindOrderById(ctx, q.OrderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
