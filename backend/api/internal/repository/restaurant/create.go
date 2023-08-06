package restaurant

import (
	"backend/api/internal/models"
	"backend/api/pkg/constants"
	"context"
	"log"
)

// Create: Create a single restaurant
func (repo *RestaurantRepository) Create(input models.Restaurant) (models.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Insert the return id
	query := constants.CreateRestaurant
	var output models.Restaurant

	row := repo.db.QueryRowContext(ctx, query,
		input.OwnerId,
		input.Name,
		input.Address,
		input.CreatedAt,
		input.UpdatedAt,
	)

	log.Println("CreateRestaurant: ", query,
		input.OwnerId,
		input.Name,
		input.Address,
		input.CreatedAt,
		input.UpdatedAt)

	err := row.Scan(
		&output.ID,
	)

	if err != nil {
		return models.Restaurant{}, err
	}

	// Get inserted restaurant to return
	queryGetRestaurant := constants.GetRestaurantByID

	row = repo.db.QueryRowContext(ctx, queryGetRestaurant, output.ID)

	log.Println("GetRestaurantByID after CreateRestaurant: ", query, output.ID)

	err = row.Scan(
		&output.ID,
		&output.OwnerId,
		&output.Name,
		&output.Address,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		return output, err
	}

	return output, nil
}
