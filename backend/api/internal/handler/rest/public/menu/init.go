package menu

import "backend/api/internal/controller/menu"

// IdRequestPayload: read id from request payload
type IdRequestPayload struct {
	Id int `json:"id"`
}

// CreateMenuRequestPayload: read creating Menu from request payload
type CreateMenuRequestPayload struct {
	RestaurantId int    `json:"restaurant_id"`
	Name         string `json:"name"`
}

// UpdateMenuRequestPayload: read updating Menu from request payload
type UpdateMenuRequestPayload struct {
	RestaurantId int    `json:"restaurant_id"`
	Name         string `json:"name"`
}

// MenuHandler: Menu handler
type MenuHandler struct {
	controller menu.MenuController
}

// NewMenuHandler: Create new Menu handler
func NewMenuHandler(handler menu.MenuController) *MenuHandler {
	return &MenuHandler{
		controller: handler,
	}
}
