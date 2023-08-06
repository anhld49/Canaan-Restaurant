package models

import (
	"time"
)

// Menu: responsible for represent Menu entity
type Menu struct {
	ID           int
	RestaurantId int
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
