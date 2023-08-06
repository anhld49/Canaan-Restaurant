package menu

import (
	"context"
	"errors"
	"log"

	"backend/api/internal/models"
	"backend/api/pkg/constants"
)

// Delete: Delete a single menu
func (repo *MenuRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Get original menu
	queryGetMenu := constants.GetMenuByID
	var original models.Menu
	row := repo.db.QueryRowContext(ctx, queryGetMenu, id)

	log.Println("GetMenuByID BEFORE UpdateMenu: ", queryGetMenu, id)

	// Fill nil value in input by original menu data
	err := row.Scan(
		&original.ID,
		&original.RestaurantId,
		&original.Name,
		&original.CreatedAt,
		&original.UpdatedAt,
	)

	if err != nil || original.ID == 0 {
		return errors.New(constants.UNKNOWN_RESTAURANT)
	}

	query := constants.DeleteMenu

	_, err = repo.db.ExecContext(ctx, query,
		id,
	)

	log.Println("DeleteMenu: ", query, id)

	if err != nil {
		return err
	}

	return nil
}
