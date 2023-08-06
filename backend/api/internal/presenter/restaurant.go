package presenter

import "time"

// Restaurant: responsible for formatting Restaurant generated as a response
type Restaurant struct {
	ID        int       `json:"id"`
	OwnerId   int       `json:"owner_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type RestaurantResponse struct {
	Success    bool       `json:"success"`
	Message    string     `json:"message"`
	Restaurant Restaurant `json:"restaurant"`
}

type DeleteRestaurantResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
