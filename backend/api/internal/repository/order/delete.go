package order

import (
	"context"
	"errors"
	"log"

	"backend/api/internal/models"
	"backend/api/pkg/constants"
)

// Delete: Delete a single order
func (repo *OrderRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Get original order
	queryGetOrder := constants.GetOrderByID
	var original models.Order
	row := repo.db.QueryRowContext(ctx, queryGetOrder, id)

	log.Println("GetOrderByID BEFORE DeleteOrder: ", queryGetOrder, id)

	// Fill nil value in input by original order data
	err := row.Scan(
		&original.ID,
		&original.UserId,
		&original.DriverId,
		&original.Amount,
		&original.CreatedAt,
		&original.UpdatedAt,
	)

	if err != nil || original.ID == 0 || original.UserId == 0 || original.DriverId == 0 {
		return errors.New(constants.UNKNOWN_DISH)
	}

	query := constants.DeleteOrder

	_, err = repo.db.ExecContext(ctx, query,
		id,
	)

	log.Println("DeleteOrder: ", query, id)

	if err != nil {
		return err
	}

	return nil
}
