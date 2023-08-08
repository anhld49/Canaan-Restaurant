package mod

import "time"

type OrderDish struct {
	DishId   int
	Quantity int
}

type Order struct {
	ID          int
	UserId      int
	DriverId    int
	Amount      float64
	OrderDishes []OrderDish
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
