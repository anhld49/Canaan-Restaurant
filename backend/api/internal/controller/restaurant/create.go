package restaurant

import "backend/api/internal/models"

// Create: Create single restaurant by ID
func (c RestaurantController) Create(restaurant models.Restaurant) (models.Restaurant, error) {
	data, err := c.restaurantRepo.Create(restaurant)

	if err != nil {
		return data, err
	}

	return data, nil
}
