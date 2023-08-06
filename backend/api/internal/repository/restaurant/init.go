package restaurant

import (
	"backend/api/internal/models"
	"database/sql"
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

// RestaurantRepositoryInterface interface
type RestaurantRepositoryInterface interface {
	Reader
	Writer
}

// RestaurantRepository: Restaurant Repository
type RestaurantRepository struct {
	db *sql.DB
}

// NewRestaurantRepo: create new Restaurant repository
func NewRestaurantRepo(db *sql.DB) *RestaurantRepository {
	return &RestaurantRepository{
		db: db,
	}
}
