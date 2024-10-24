package api

import (
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Domain/Model"
	"DDD-OrderingSystem/Infrastructure/Repository"
	"encoding/json"
	"net/http"
)

// OrderController 订单控制器
type OrderController struct {
	orderRepository Repository.OrderRepository
	commandHandler  Command.Handler
}

// NewOrderController 创建订单控制器实例
func NewOrderController(orderRepository Repository.OrderRepository, commandHandler Command.Handler) *OrderController {
	return &OrderController{
		orderRepository: orderRepository,
		commandHandler:  commandHandler,
	}
}

// CreateOrder 创建订单
func (oc *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// 解析请求体
	var order Model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 创建命令
	createOrderCommand := Command.NewOrderCreateCommand(order.CustomerID, order.Items)

	// 处理命令
	if err := oc.commandHandler.Handle(context.Background(), createOrderCommand); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createOrderCommand.GetOrder())
}

// GetOrder 获取订单
func (oc *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	order, err := oc.orderRepository.FindById(context.Background(), orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
