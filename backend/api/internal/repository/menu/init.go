package menu

import (
	"backend/api/internal/models"
	"database/sql"
)

// Reader interface
type Reader interface {
	List() ([]models.Menu, error)
	Get(id int) (models.Menu, error)
}

// Writer user writer
type Writer interface {
	Create(models.Menu) (models.Menu, error)
	Update(models.Menu) (models.Menu, error)
	Delete(id int) error
}

// MenuRepositoryInterface interface
type MenuRepositoryInterface interface {
	Reader
	Writer
}

// MenuRepository: Menu Repository
type MenuRepository struct {
	db *sql.DB
}

// NewMenuRepo: create new Menu repository
func NewMenuRepo(db *sql.DB) *MenuRepository {
	return &MenuRepository{
		db: db,
	}
}
