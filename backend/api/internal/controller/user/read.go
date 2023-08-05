package user

import (
	"backend/api/internal/models"
)

// List: Get all users
func (c UserController) List() ([]models.User, error) {
	data, err := c.userRepo.List()

	if err != nil {
		return nil, err
	}

	var users []models.User
	for _, d := range data {
		users = append(users, models.User{
			ID:        d.ID,
			FirstName: d.FirstName,
			LastName:  d.LastName,
			Email:     d.Email,
			Password:  d.Password,
			Role:      d.Role,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return users, nil
}

// Get: Get single user by email
func (c UserController) Get(email string) (models.User, error) {
	data, err := c.userRepo.Get(email)

	if err != nil {
		return data, err
	}

	return data, nil
}

// Get: Get single user by ID
func (c UserController) GetByID(id int) (models.User, error) {
	data, err := c.userRepo.GetByID(id)

	if err != nil {
		return data, err
	}

	return data, nil
}
