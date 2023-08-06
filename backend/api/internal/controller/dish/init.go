package dish

import (
	"backend/api/internal/models"
	"backend/api/internal/repository/dish"
)

// Reader interface
type Reader interface {
	List() ([]models.Dish, error)
	Get(id int) (models.Dish, error)
}

// Writer user writer
type Writer interface {
	Create(models.Dish) (models.Dish, error)
	Update(models.Dish) (models.Dish, error)
	Delete(id int) error
}

// DishControllerInterface interface
type DishControllerInterface interface {
	Reader
	Writer
}

// DishController: Dish Controller
type DishController struct {
	dishRepo dish.DishRepository
}

// NewDishController: Create new Dish Controller
func NewDishController(r dish.DishRepository) *DishController {
	return &DishController{
		dishRepo: r,
	}
}
