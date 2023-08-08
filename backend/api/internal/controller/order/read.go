package order

import "backend/api/internal/mod"

// List: Get all orders
func (c OrderController) List(userId int) ([]mod.Order, error) {
	data, err := c.orderRepo.List(userId)

	if err != nil {
		return nil, err
	}

	var orders []mod.Order
	for _, order := range data {
		orders = append(orders, mod.Order{
			ID:          order.ID,
			DriverId:    order.DriverId,
			Amount:      order.Amount,
			OrderDishes: order.OrderDishes,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		})
	}

	return orders, nil
}

// Get: Get single order by ID
func (c OrderController) Get(id int) (mod.Order, error) {
	data, err := c.orderRepo.Get(id)

	if err != nil {
		return data, err
	}

	return data, nil
}
