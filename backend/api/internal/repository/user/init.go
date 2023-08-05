package user

import (
	"backend/api/internal/models"
	"database/sql"
)

// Repository interface
type Repository interface {
	List() ([]models.User, error)
	Get(email string) (models.User, error)
	GetByID(id int) (models.User, error)
}

// UserRepository: User Repository
type UserRepository struct {
	db *sql.DB
}

// NewUserRepo: create new User repository
func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
