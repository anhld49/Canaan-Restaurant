package dish

import (
	"backend/api/internal/models"
	"backend/api/pkg/constants"
	"context"
	"log"
)

// Update: Update a single dish
func (repo *DishRepository) Update(input models.Dish) (models.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Get original dish
	queryGetDish := constants.GetDishByID
	var original models.Dish
	row := repo.db.QueryRowContext(ctx, queryGetDish, input.ID)

	log.Println("GetDishByID BEFORE UpdateDish: ", queryGetDish, input.ID)

	// Fill nil value in input by original dish data
	err := row.Scan(
		&original.ID,
		&original.MenuId,
		&original.Name,
		&original.Price,
		&original.CreatedAt,
		&original.UpdatedAt,
	)

	if err != nil {
		return original, err
	}

	if input.MenuId == 0 {
		input.MenuId = original.MenuId
	}
	if input.Name == "" {
		input.Name = original.Name
	}

	// Update Dish
	query := constants.UpdateDish
	var output models.Dish

	row = repo.db.QueryRowContext(ctx, query,
		input.MenuId,
		input.Name,
		input.Price,
		input.UpdatedAt,
		input.ID,
	)

	log.Println("UpdateDish: ", query,
		input.MenuId,
		input.Name,
		input.Price,
		input.UpdatedAt,
		input.ID,
	)

	err = row.Scan(
		&output.ID,
	)

	if err != nil {
		return models.Dish{}, err
	}

	// Get updated dish to return
	row = repo.db.QueryRowContext(ctx, queryGetDish, output.ID)

	log.Println("GetDishByID AFTER UpdateDish: ", queryGetDish, output.ID)

	// Fill nil value in input by original dish data
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
