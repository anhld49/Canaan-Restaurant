package restaurant

import (
	"context"
	"errors"
	"log"

	"backend/api/internal/models"
	"backend/api/pkg/constants"
)

// Delete: Delete a single restaurant
func (repo *RestaurantRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Get original restaurant
	queryGetRestaurant := constants.GetRestaurantByID
	var original models.Restaurant
	row := repo.db.QueryRowContext(ctx, queryGetRestaurant, id)

	log.Println("GetRestaurantByID BEFORE UpdateRestaurant: ", queryGetRestaurant, id)

	// Fill nil value in input by original restaurant data
	err := row.Scan(
		&original.ID,
		&original.OwnerId,
		&original.Name,
		&original.Address,
		&original.CreatedAt,
		&original.UpdatedAt,
	)

	if err != nil || original.ID == 0 {
		return errors.New(constants.UNKNOWN_RESTAURANT)
	}

	query := constants.DeleteRestaurant

	_, err = repo.db.ExecContext(ctx, query,
		id,
	)

	log.Println("DeleteRestaurant: ", query, id)

	if err != nil {
		return err
	}

	return nil
}
