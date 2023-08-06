package models

import (
	"time"
)

// Dish: responsible for represent Dish entity
type Dish struct {
	ID        int
	MenuId    int
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
