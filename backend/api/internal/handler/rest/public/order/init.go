package order

import "backend/api/internal/controller/order"

// IdRequestPayload: read id from request payload
type IdRequestPayload struct {
	Id int `json:"id"`
}

type OrderDishRequestPayload struct {
	DishId   int `json:"dish_id"`
	Quantity int `json:"quantity"`
}

// CreateOrderRequestPayload: read creating Order from request payload
type CreateOrderRequestPayload struct {
	DriverId  int                       `json:"driver_id"`
	OrderDish []OrderDishRequestPayload `json:"order_dish"`
}

// OrderHandler: Order handler
type OrderHandler struct {
	controller order.OrderController
}

// NewOrderHandler: Create new Order handler
func NewOrderHandler(handler order.OrderController) *OrderHandler {
	return &OrderHandler{
		controller: handler,
	}
}
