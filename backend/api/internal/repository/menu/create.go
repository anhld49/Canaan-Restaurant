package menu

import (
	"backend/api/internal/models"
	"backend/api/pkg/constants"
	"context"
	"log"
)

// Create: Create a single menu
func (repo *MenuRepository) Create(input models.Menu) (models.Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Insert the return id
	query := constants.CreateMenu
	var output models.Menu

	row := repo.db.QueryRowContext(ctx, query,
		input.RestaurantId,
		input.Name,
		input.CreatedAt,
		input.UpdatedAt,
	)

	log.Println("CreateMenu: ", query,
		input.RestaurantId,
		input.Name,
		input.CreatedAt,
		input.UpdatedAt)

	err := row.Scan(
		&output.ID,
	)

	if err != nil {
		return models.Menu{}, err
	}

	// Get inserted menu to return
	queryGetMenu := constants.GetMenuByID

	row = repo.db.QueryRowContext(ctx, queryGetMenu, output.ID)

	log.Println("GetMenuByID after CreateMenu: ", query, output.ID)

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
