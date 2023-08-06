package models

import (
	"time"
)

// Order: responsible for represent Order entity
type Order struct {
	ID        int
	MenuId    int
	DriverId  int
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
