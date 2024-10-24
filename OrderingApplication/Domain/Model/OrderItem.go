package model

// OrderItem 订单项模型
type OrderItem struct {
	ID        uint64
	ProductID uint64
	Quantity  int
	Price     float64
}

// NewOrderItem 创建新的订单项
func NewOrderItem(productID uint64, quantity int, price float64) *OrderItem {
	return &OrderItem{
		ID:        0, // 实际应用中应由数据库生成
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
}
