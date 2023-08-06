package restaurant

import "backend/api/internal/controller/restaurant"

// IdRequestPayload: read id from request payload
type IdRequestPayload struct {
	Id int `json:"id"`
}

// CreateRestaurantRequestPayload: read creating restaurant from request payload
type CreateRestaurantRequestPayload struct {
	OwnerId int    `json:"owner_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// UpdateRestaurantRequestPayload: read updating restaurant from request payload
type UpdateRestaurantRequestPayload struct {
	OwnerId int    `json:"owner_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// RestaurantHandler: Restaurant handler
type RestaurantHandler struct {
	controller restaurant.RestaurantController
}

// NewRestaurantHandler: Create new restaurant handler
func NewRestaurantHandler(handler restaurant.RestaurantController) *RestaurantHandler {
	return &RestaurantHandler{
		controller: handler,
	}
}
