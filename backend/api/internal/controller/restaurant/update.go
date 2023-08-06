package restaurant

import "backend/api/internal/models"

// Update: Update single restaurant by ID
func (c RestaurantController) Update(restaurant models.Restaurant) (models.Restaurant, error) {
	data, err := c.restaurantRepo.Update(restaurant)

	if err != nil {
		return data, err
	}

	return data, nil
}
