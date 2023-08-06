package restaurant

import (
	"backend/api/internal/models"
	"backend/api/pkg/constants"
	"context"
	"log"
)

// Update: Update a single restaurant
func (repo *RestaurantRepository) Update(input models.Restaurant) (models.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Get original restaurant
	queryGetRestaurant := constants.GetRestaurantByID
	var original models.Restaurant
	row := repo.db.QueryRowContext(ctx, queryGetRestaurant, input.ID)

	log.Println("GetRestaurantByID BEFORE UpdateRestaurant: ", queryGetRestaurant, input.ID)

	// Fill nil value in input by original restaurant data
	err := row.Scan(
		&original.ID,
		&original.OwnerId,
		&original.Name,
		&original.Address,
		&original.CreatedAt,
		&original.UpdatedAt,
	)

	if err != nil {
		return original, err
	}

	if input.OwnerId == 0 {
		input.OwnerId = original.OwnerId
	}
	if input.Name == "" {
		input.Name = original.Name
	}
	if input.Address == "" {
		input.Address = original.Address
	}

	// Update Restaurant
	query := constants.UpdateRestaurant
	var output models.Restaurant

	row = repo.db.QueryRowContext(ctx, query,
		input.OwnerId,
		input.Name,
		input.Address,
		input.UpdatedAt,
		input.ID,
	)

	log.Println("UpdateRestaurant: ", query,
		input.OwnerId,
		input.Name,
		input.Address,
		input.UpdatedAt,
		input.ID)

	err = row.Scan(
		&output.ID,
	)

	if err != nil {
		return models.Restaurant{}, err
	}

	// Get updated restaurant to return
	row = repo.db.QueryRowContext(ctx, queryGetRestaurant, output.ID)

	log.Println("GetRestaurantByID AFTER UpdateRestaurant: ", queryGetRestaurant, output.ID)

	// Fill nil value in input by original restaurant data
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
