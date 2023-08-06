package restaurant

import "backend/api/internal/controller/restaurant"

// IdRequestPayload: read id from request payload
type IdRequestPayload struct {
	Id int `json:"id"`
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
