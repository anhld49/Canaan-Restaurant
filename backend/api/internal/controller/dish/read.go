package dish

import (
	"backend/api/internal/models"
)

// List: Get all dishs
func (c DishController) List() ([]models.Dish, error) {
	data, err := c.dishRepo.List()

	if err != nil {
		return nil, err
	}

	var dishs []models.Dish
	for _, d := range data {
		dishs = append(dishs, models.Dish{
			ID:        d.ID,
			MenuId:    d.MenuId,
			Name:      d.Name,
			Price:     d.Price,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return dishs, nil
}

// Get: Get single dish by ID
func (c DishController) Get(id int) (models.Dish, error) {
	data, err := c.dishRepo.Get(id)

	if err != nil {
		return data, err
	}

	return data, nil
}

// GetDishesByMenuId: Get Dishes By Menu Id
func (c DishController) GetDishesByMenuId(id int) ([]models.Dish, error) {
	data, err := c.dishRepo.GetDishesByMenuId(id)

	if err != nil {
		return data, err
	}

	return data, nil
}
