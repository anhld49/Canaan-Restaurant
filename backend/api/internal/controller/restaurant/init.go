package restaurant

import (
	"backend/api/internal/models"
	"backend/api/internal/repository/restaurant"
)

// Reader interface
type Reader interface {
	List() ([]models.Restaurant, error)
	Get(id int) (models.Restaurant, error)
}

// Writer user writer
// type Writer interface {
// 	CreateFriendship(email string, friend string) (mod.UserResponse, error)
// 	CreateSubscribe(requestor string, target string) (mod.UserResponse, error)
// 	CreateBlock(requestor string, target string) error
// }

// RestaurantControllerInterface interface
type RestaurantControllerInterface interface {
	Reader
	// Writer
}

// RestaurantController: Restaurant Controller
type RestaurantController struct {
	restaurantRepo restaurant.Repository
}

// NewRestaurantController: Create new Restaurant Controller
func NewRestaurantController(r restaurant.Repository) *RestaurantController {
	return &RestaurantController{
		restaurantRepo: r,
	}
}
