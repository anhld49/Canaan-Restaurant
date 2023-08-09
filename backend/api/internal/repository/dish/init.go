package dish

import (
	"backend/api/internal/models"
	"database/sql"
)

// Reader interface
type Reader interface {
	List() ([]models.Dish, error)
	Get(id int) (models.Dish, error)
	GetDishesByMenuId(id int) ([]models.Dish, error)
}

// Writer user writer
type Writer interface {
	Create(models.Dish) (models.Dish, error)
	Update(models.Dish) (models.Dish, error)
	Delete(id int) error
}

// DishRepositoryInterface interface
type DishRepositoryInterface interface {
	Reader
	Writer
}

// DishRepository: Dish Repository
type DishRepository struct {
	db *sql.DB
}

// NewDishRepo: create new Dish repository
func NewDishRepo(db *sql.DB) *DishRepository {
	return &DishRepository{
		db: db,
	}
}
