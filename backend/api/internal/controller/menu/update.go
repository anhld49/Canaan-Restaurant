package menu

import "backend/api/internal/models"

// Update: Update single menu by ID
func (c MenuController) Update(menu models.Menu) (models.Menu, error) {
	data, err := c.menuRepo.Update(menu)

	if err != nil {
		return data, err
	}

	return data, nil
}
