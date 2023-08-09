package dish

import (
	"context"
	"log"
	"time"

	"backend/api/internal/models"
	"backend/api/pkg/constants"
)

const dbTimeout = time.Second * 3

// List: Get All dishs from database
func (repo *DishRepository) List() ([]models.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetAllDishs

	rows, err := repo.db.QueryContext(ctx, query)

	log.Println("GetAllDishs: ", query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishs []models.Dish

	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(
			&dish.ID,
			&dish.MenuId,
			&dish.Name,
			&dish.Price,
			&dish.CreatedAt,
			&dish.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		dishs = append(dishs, dish)
	}

	return dishs, nil
}

// Get: Get a single dish by id
func (repo *DishRepository) Get(id int) (models.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetDishByID

	dish := models.Dish{}

	row := repo.db.QueryRowContext(ctx, query, id)

	log.Println("GetDishByID: ", query, id)

	err := row.Scan(
		&dish.ID,
		&dish.MenuId,
		&dish.Name,
		&dish.Price,
		&dish.CreatedAt,
		&dish.UpdatedAt,
	)

	if err != nil {
		return dish, err
	}

	return dish, nil

}

// GetDishesByMenuId: Get Dishes By Menu Id
func (repo *DishRepository) GetDishesByMenuId(id int) ([]models.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetDishesByMenuId

	rows, err := repo.db.QueryContext(ctx, query, id)

	log.Println("GetDishesByMenuId: ", query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishs []models.Dish

	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(
			&dish.ID,
			&dish.MenuId,
			&dish.Name,
			&dish.Price,
			&dish.CreatedAt,
			&dish.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		dishs = append(dishs, dish)
	}

	return dishs, nil

}
