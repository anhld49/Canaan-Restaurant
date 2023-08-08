package order

import (
	"backend/api/internal/mod"
	"backend/api/internal/repository/order"
)

// Reader interface
type Reader interface {
	List(userId int) ([]mod.Order, error)
	Get(id int) (mod.Order, error)
}

// Writer user writer
type Writer interface {
	Create(mod.Order) (mod.Order, error)
	Delete(id int) error
}

// OrderControllerInterface interface
type OrderControllerInterface interface {
	Reader
	Writer
}

// OrderController: Order Controller
type OrderController struct {
	orderRepo order.OrderRepository
}

// NewOrderController: Create new Order Controller
func NewOrderController(r order.OrderRepository) *OrderController {
	return &OrderController{
		orderRepo: r,
	}
}
