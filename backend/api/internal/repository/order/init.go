package order

import (
	"backend/api/internal/mod"
	"database/sql"
)

// Reader interface
type Reader interface {
	List(userId int) ([]mod.Order, error)
	Get(id int) (mod.Order, error)
}

// Writer user writer
type Writer interface {
	Create(mod.Order) (mod.Order, error)
	Delete(id int) error
}

// OrderRepositoryInterface interface
type OrderRepositoryInterface interface {
	Reader
	Writer
}

// OrderRepository: Order Repository
type OrderRepository struct {
	db *sql.DB
}

// NewOrderRepo: create new Order repository
func NewOrderRepo(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}
