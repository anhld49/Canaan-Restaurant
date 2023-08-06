package presenter

import "time"

// Menu: responsible for formatting Menu generated as a response
type Menu struct {
	ID           int       `json:"id"`
	RestaurantId int       `json:"restaurant_id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type MenuResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Menu    Menu   `json:"menu"`
}

type DeleteMenuResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
