package dish

import (
	"backend/api/internal/models"
	"backend/api/pkg/constants"
	"context"
	"log"
)

// Create: Create a single dish
func (repo *DishRepository) Create(input models.Dish) (models.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Insert the return id
	query := constants.CreateDish
	var output models.Dish

	row := repo.db.QueryRowContext(ctx, query,
		input.MenuId,
		input.Name,
		input.Price,
		input.CreatedAt,
		input.UpdatedAt,
	)

	log.Println("CreateDish: ", query,
		input.MenuId,
		input.Name,
		input.Price,
		input.CreatedAt,
		input.UpdatedAt,
	)

	err := row.Scan(
		&output.ID,
	)

	if err != nil {
		return models.Dish{}, err
	}

	// Get inserted dish to return
	queryGetDish := constants.GetDishByID

	row = repo.db.QueryRowContext(ctx, queryGetDish, output.ID)

	log.Println("GetDishByID after CreateDish: ", query, output.ID)

	err = row.Scan(
		&output.ID,
		&output.MenuId,
		&output.Name,
		&output.Price,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		return output, err
	}

	return output, nil
}
