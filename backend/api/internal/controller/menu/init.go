package menu

import (
	"backend/api/internal/models"
	"backend/api/internal/repository/menu"
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

// MenuControllerInterface interface
type MenuControllerInterface interface {
	Reader
	Writer
}

// MenuController: Menu Controller
type MenuController struct {
	menuRepo menu.MenuRepository
}

// NewMenuController: Create new Menu Controller
func NewMenuController(r menu.MenuRepository) *MenuController {
	return &MenuController{
		menuRepo: r,
	}
}
