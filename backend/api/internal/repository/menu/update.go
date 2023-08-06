package menu

import (
	"backend/api/internal/models"
	"backend/api/pkg/constants"
	"context"
	"log"
)

// Update: Update a single menu
func (repo *MenuRepository) Update(input models.Menu) (models.Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Get original menu
	queryGetMenu := constants.GetMenuByID
	var original models.Menu
	row := repo.db.QueryRowContext(ctx, queryGetMenu, input.ID)

	log.Println("GetMenuByID BEFORE UpdateMenu: ", queryGetMenu, input.ID)

	// Fill nil value in input by original menu data
	err := row.Scan(
		&original.ID,
		&original.RestaurantId,
		&original.Name,
		&original.CreatedAt,
		&original.UpdatedAt,
	)

	if err != nil {
		return original, err
	}

	if input.RestaurantId == 0 {
		input.RestaurantId = original.RestaurantId
	}
	if input.Name == "" {
		input.Name = original.Name
	}

	// Update Menu
	query := constants.UpdateMenu
	var output models.Menu

	row = repo.db.QueryRowContext(ctx, query,
		input.RestaurantId,
		input.Name,
		input.UpdatedAt,
		input.ID,
	)

	log.Println("UpdateMenu: ", query,
		input.RestaurantId,
		input.Name,
		input.UpdatedAt,
		input.ID)

	err = row.Scan(
		&output.ID,
	)

	if err != nil {
		return models.Menu{}, err
	}

	// Get updated menu to return
	row = repo.db.QueryRowContext(ctx, queryGetMenu, output.ID)

	log.Println("GetMenuByID AFTER UpdateMenu: ", queryGetMenu, output.ID)

	// Fill nil value in input by original menu data
	err = row.Scan(
		&output.ID,
		&output.RestaurantId,
		&output.Name,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		return output, err
	}

	return output, nil

}
