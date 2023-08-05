package user

import (
	"context"
	"log"
	"time"

	"backend/api/internal/models"
	"backend/api/pkg/constants"
)

const dbTimeout = time.Second * 3

// List: Get All users from database
func (repo *UserRepository) List() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetAllUsers

	rows, err := repo.db.QueryContext(ctx, query)

	log.Println("GetAllUsers: ", query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Get: Get a single user by email
func (repo *UserRepository) Get(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetUser

	user := models.User{}

	row := repo.db.QueryRowContext(ctx, query, email)

	log.Println("GetUser: ", query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil

}

// GetByID: Get a single user by id
func (repo *UserRepository) GetByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetUserByID

	user := models.User{}

	row := repo.db.QueryRowContext(ctx, query, id)

	log.Println("GetUserByID: ", query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil

}
