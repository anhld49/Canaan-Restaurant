package menu

import (
	"backend/api/internal/models"
)

// List: Get all menus
func (c MenuController) List() ([]models.Menu, error) {
	data, err := c.menuRepo.List()

	if err != nil {
		return nil, err
	}

	var menus []models.Menu
	for _, d := range data {
		menus = append(menus, models.Menu{
			ID:           d.ID,
			RestaurantId: d.RestaurantId,
			Name:         d.Name,
			CreatedAt:    d.CreatedAt,
			UpdatedAt:    d.UpdatedAt,
		})
	}

	return menus, nil
}

// Get: Get single menu by ID
func (c MenuController) Get(id int) (models.Menu, error) {
	data, err := c.menuRepo.Get(id)

	if err != nil {
		return data, err
	}

	return data, nil
}
