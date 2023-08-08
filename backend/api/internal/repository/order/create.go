package order

import (
	"backend/api/internal/mod"
	"backend/api/pkg/constants"
	"context"
	"log"
)

// Create: Create a single order
func (repo *OrderRepository) Create(input mod.Order) (mod.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	amount := 0.00

	// Insert Order then return id
	// Input includes: UserId, DriverId, []OrderDishes, CreatedAt, UpdatedAt
	query := constants.CreateOrderInit
	var output mod.Order

	row := repo.db.QueryRowContext(ctx, query,
		input.UserId,
		input.DriverId,
		input.CreatedAt,
		input.UpdatedAt,
	)

	log.Println("CreateOrderInit: ", query,
		input.UserId,
		input.DriverId,
		input.CreatedAt,
		input.UpdatedAt,
	)

	err := row.Scan(
		&output.ID,
	)

	if err != nil {
		return mod.Order{}, err
	}

	// Insert into OrderDish and calculate Amount
	query = constants.CreateOrderDish

	for _, v := range input.OrderDishes {
		repo.db.QueryRowContext(ctx, query,
			output.ID,
			v.DishId,
			v.Quantity,
			input.CreatedAt,
			input.UpdatedAt,
		)

		log.Println("CreateOrderDish: ", query,
			output.ID,
			v.DishId,
			v.Quantity,
			input.CreatedAt,
			input.UpdatedAt,
		)

		query2 := constants.GetDishPriceByDishId
		row = repo.db.QueryRowContext(ctx, query2, v.DishId)
		price := 0.00

		log.Println("GetDishPriceByDishId: ", query2, v.DishId)

		err := row.Scan(&price)

		if err != nil {
			return mod.Order{}, err
		}
		amount += price * float64(v.Quantity)
	}

	// Insert Amount into Order
	query = constants.InsertAmountIntoOrder

	repo.db.QueryRowContext(ctx, query,
		amount,
		output.ID,
	)

	log.Println("InsertAmountIntoOrder: ", query,
		amount,
		output.ID,
	)

	// Get inserted order to return
	queryGetOrder := constants.GetOrderByID

	row = repo.db.QueryRowContext(ctx, queryGetOrder, output.ID)

	log.Println("GetOrderByID", query, output.ID)

	err = row.Scan(
		&output.ID,
		&output.UserId,
		&output.DriverId,
		&output.Amount,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		return mod.Order{}, err
	}

	// Get OrderDish to return
	query = constants.GetOrderDishByOrderId
	rows, err := repo.db.QueryContext(ctx, query, output.ID)
	log.Println("Create Order, GetOrderDishByOrderId: ", query, output.ID)

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

		output.OrderDishes = append(output.OrderDishes, orderDish)
	}

	return output, nil
}
