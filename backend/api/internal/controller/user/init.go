package user

import (
	"backend/api/internal/models"
	"backend/api/internal/repository/user"
)

// Reader interface
type Reader interface {
	List() ([]models.User, error)
	Get(email string) (models.User, error)
	GetByID(id int) (models.User, error)
}

// Writer user writer
// type Writer interface {
// 	CreateFriendship(email string, friend string) (mod.UserResponse, error)
// 	CreateSubscribe(requestor string, target string) (mod.UserResponse, error)
// 	CreateBlock(requestor string, target string) error
// }

// UserControllerInterface interface
type UserControllerInterface interface {
	Reader
	// Writer
}

// UserController: User Controller
type UserController struct {
	userRepo user.Repository
}

// NewUserController: Create new User Controller
func NewUserController(r user.Repository) *UserController {
	return &UserController{
		userRepo: r,
	}
}
