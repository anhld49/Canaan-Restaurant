package presenter

import "time"

type OrderDish struct {
	DishId   int `json:"dish_id"`
	Quantity int `json:"quantity"`
}

// Order: responsible for formatting Order generated as a response
type Order struct {
	ID          int         `json:"id"`
	DriverId    int         `json:"driver_id"`
	Amount      float64     `json:"amount"`
	OrderDishes []OrderDish `json:"order_dishes"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
}

type OrderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Order   Order  `json:"order"`
}

type DeleteOrderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
