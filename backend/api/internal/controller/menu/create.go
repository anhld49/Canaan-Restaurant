package menu

import "backend/api/internal/models"

// Create: Create single menu by ID
func (c MenuController) Create(menu models.Menu) (models.Menu, error) {
	data, err := c.menuRepo.Create(menu)

	if err != nil {
		return data, err
	}

	return data, nil
}
