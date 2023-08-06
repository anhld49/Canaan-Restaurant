package dish

import "backend/api/internal/models"

// Update: Update single dish by ID
func (c DishController) Update(dish models.Dish) (models.Dish, error) {
	data, err := c.dishRepo.Update(dish)

	if err != nil {
		return data, err
	}

	return data, nil
}
