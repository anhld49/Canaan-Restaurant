package order

import (
	"context"
	"errors"
	"log"
	"time"

	"backend/api/internal/mod"
	"backend/api/pkg/constants"
)

const dbTimeout = time.Second * 3

// List: Get All orders from database
func (repo *OrderRepository) List(userId int) ([]mod.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Get Orders by UserId
	query := constants.GetAllOrders

	rows, err := repo.db.QueryContext(ctx, query, userId)

	log.Println("GetAllOrders: ", query, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []mod.Order

	for rows.Next() {
		var order mod.Order
		err := rows.Scan(
			&order.ID,
			&order.UserId,
			&order.DriverId,
			&order.Amount,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	// Get All OrderDishes By OrderId
	query = constants.GetOrderDishByOrderId

	for i, order := range orders {
		rows, err = repo.db.QueryContext(ctx, query, order.ID)
		log.Println("List All Orders, GetOrderDishByOrderId: ", query, order.ID)

		if err != nil {
			return []mod.Order{}, err
		}

		for rows.Next() {
			orderDish := mod.OrderDish{}
			err := rows.Scan(
				&orderDish.DishId,
				&orderDish.Quantity,
			)
			if err != nil {
				return []mod.Order{}, err
			}

			orders[i].OrderDishes = append(orders[i].OrderDishes, mod.OrderDish{
				DishId:   orderDish.DishId,
				Quantity: orderDish.Quantity,
			})
		}
	}

	return orders, nil
}

// Get: Get a single order by id
func (repo *OrderRepository) Get(id int) (mod.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	order := mod.Order{}

	// Get Orders by UserId
	query := constants.GetOrderByID

	row := repo.db.QueryRowContext(ctx, query, id)

	log.Println("GetOrderByID: ", query, id)

	err := row.Scan(
		&order.ID,
		&order.UserId,
		&order.DriverId,
		&order.Amount,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return mod.Order{}, errors.New(constants.UNKNOWN_ORDER)
	}

	// Get All OrderDishes By OrderId
	query = constants.GetOrderDishByOrderId

	rows, err := repo.db.QueryContext(ctx, query, order.ID)
	log.Println("List All Orders, GetOrderDishByOrderId: ", query, id)

	if err != nil {
		return mod.Order{}, err
	}

	for rows.Next() {
		orderDish := mod.OrderDish{}
		err := rows.Scan(
			&orderDish.DishId,
			&orderDish.Quantity,
		)
		if err != nil {
			return mod.Order{}, err
		}

		order.OrderDishes = append(order.OrderDishes, mod.OrderDish{
			DishId:   orderDish.DishId,
			Quantity: orderDish.Quantity,
		})
	}

	return order, nil

}
