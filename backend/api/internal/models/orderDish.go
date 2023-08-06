package models

import (
	"time"
)

// OrderDish: responsible for represent OrderDish entity
type OrderDish struct {
	OrderId   int
	DishId    int
	Quantity  float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
