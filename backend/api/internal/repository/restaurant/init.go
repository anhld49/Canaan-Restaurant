package restaurant

import (
	"backend/api/internal/models"
	"database/sql"
)

// Repository interface
type Repository interface {
	List() ([]models.Restaurant, error)
	Get(id int) (models.Restaurant, error)
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
