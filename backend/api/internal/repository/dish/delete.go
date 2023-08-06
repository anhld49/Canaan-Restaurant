package dish

import (
	"context"
	"errors"
	"log"

	"backend/api/internal/models"
	"backend/api/pkg/constants"
)

// Delete: Delete a single dish
func (repo *DishRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Get original dish
	queryGetDish := constants.GetDishByID
	var original models.Dish
	row := repo.db.QueryRowContext(ctx, queryGetDish, id)

	log.Println("GetDishByID BEFORE DeleteDish: ", queryGetDish, id)

	// Fill nil value in input by original dish data
	err := row.Scan(
		&original.ID,
		&original.MenuId,
		&original.Name,
		&original.Price,
		&original.CreatedAt,
		&original.UpdatedAt,
	)

	if err != nil || original.ID == 0 {
		return errors.New(constants.UNKNOWN_DISH)
	}

	query := constants.DeleteDish

	_, err = repo.db.ExecContext(ctx, query,
		id,
	)

	log.Println("DeleteDish: ", query, id)

	if err != nil {
		return err
	}

	return nil
}
