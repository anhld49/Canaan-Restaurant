package presenter

import "time"

// Dish: responsible for formatting Dish generated as a response
type Dish struct {
	ID        int       `json:"id"`
	MenuId    int       `json:"menu_id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type DishResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Dish    Dish   `json:"dish"`
}

type DeleteDishResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
