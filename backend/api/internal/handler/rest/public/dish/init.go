package dish

import "backend/api/internal/controller/dish"

// IdRequestPayload: read id from request payload
type IdRequestPayload struct {
	Id int `json:"id"`
}

// CreateDishRequestPayload: read creating Dish from request payload
type CreateDishRequestPayload struct {
	MenuId int     `json:"menu_id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}

// UpdateDishRequestPayload: read updating Dish from request payload
type UpdateDishRequestPayload struct {
	MenuId int     `json:"menu_id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}

// DishHandler: Dish handler
type DishHandler struct {
	controller dish.DishController
}

// NewDishHandler: Create new Dish handler
func NewDishHandler(handler dish.DishController) *DishHandler {
	return &DishHandler{
		controller: handler,
	}
}
