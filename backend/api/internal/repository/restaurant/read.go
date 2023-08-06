package restaurant

import (
	"context"
	"log"
	"time"

	"backend/api/internal/models"
	"backend/api/pkg/constants"
)

const dbTimeout = time.Second * 3

// List: Get All restaurants from database
func (repo *RestaurantRepository) List() ([]models.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetAllRestaurants

	rows, err := repo.db.QueryContext(ctx, query)

	log.Println("GetAllRestaurants: ", query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []models.Restaurant

	for rows.Next() {
		var restaurant models.Restaurant
		err := rows.Scan(
			&restaurant.ID,
			&restaurant.OwnerId,
			&restaurant.Name,
			&restaurant.Address,
			&restaurant.CreatedAt,
			&restaurant.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

// Get: Get a single restaurant by id
func (repo *RestaurantRepository) Get(id int) (models.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetRestaurantByID

	restaurant := models.Restaurant{}

	row := repo.db.QueryRowContext(ctx, query, id)

	log.Println("GetRestaurantByID: ", query, id)

	err := row.Scan(
		&restaurant.ID,
		&restaurant.OwnerId,
		&restaurant.Name,
		&restaurant.Address,
		&restaurant.CreatedAt,
		&restaurant.UpdatedAt,
	)

	if err != nil {
		return restaurant, err
	}

	return restaurant, nil

}
