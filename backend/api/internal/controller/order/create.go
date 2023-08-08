package order

import "backend/api/internal/mod"

// Create: Create single order by ID
func (c OrderController) Create(order mod.Order) (mod.Order, error) {
	data, err := c.orderRepo.Create(order)

	if err != nil {
		return data, err
	}

	return data, nil
}
