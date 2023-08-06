package restaurant

import (
	"backend/api/internal/models"
)

// List: Get all restaurants
func (c RestaurantController) List() ([]models.Restaurant, error) {
	data, err := c.restaurantRepo.List()

	if err != nil {
		return nil, err
	}

	var restaurants []models.Restaurant
	for _, d := range data {
		restaurants = append(restaurants, models.Restaurant{
			ID:        d.ID,
			OwnerId:   d.OwnerId,
			Name:      d.Name,
			Address:   d.Address,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return restaurants, nil
}

// Get: Get single restaurant by ID
func (c RestaurantController) Get(id int) (models.Restaurant, error) {
	data, err := c.restaurantRepo.Get(id)

	if err != nil {
		return data, err
	}

	return data, nil
}
