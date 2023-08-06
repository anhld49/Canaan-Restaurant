package models

import (
	"time"
)

// Restaurant: responsible for represent Restaurant entity
type Restaurant struct {
	ID        int
	OwnerId   int
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
