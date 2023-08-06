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
type Writer interface {
	Create(models.Restaurant) (models.Restaurant, error)
	Update(models.Restaurant) (models.Restaurant, error)
	Delete(id int) error
}

// RestaurantControllerInterface interface
type RestaurantControllerInterface interface {
	Reader
	Writer
}

// RestaurantController: Restaurant Controller
type RestaurantController struct {
	restaurantRepo restaurant.RestaurantRepository
}

// NewRestaurantController: Create new Restaurant Controller
func NewRestaurantController(r restaurant.RestaurantRepository) *RestaurantController {
	return &RestaurantController{
		restaurantRepo: r,
	}
}
