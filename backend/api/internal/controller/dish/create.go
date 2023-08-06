package dish

import "backend/api/internal/models"

// Create: Create single dish by ID
func (c DishController) Create(dish models.Dish) (models.Dish, error) {
	data, err := c.dishRepo.Create(dish)

	if err != nil {
		return data, err
	}

	return data, nil
}
