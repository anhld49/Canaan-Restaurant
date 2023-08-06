package menu

import (
	"context"
	"log"
	"time"

	"backend/api/internal/models"
	"backend/api/pkg/constants"
)

const dbTimeout = time.Second * 3

// List: Get All menus from database
func (repo *MenuRepository) List() ([]models.Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetAllMenus

	rows, err := repo.db.QueryContext(ctx, query)

	log.Println("GetAllMenus: ", query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []models.Menu

	for rows.Next() {
		var menu models.Menu
		err := rows.Scan(
			&menu.ID,
			&menu.RestaurantId,
			&menu.Name,
			&menu.CreatedAt,
			&menu.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

// Get: Get a single menu by id
func (repo *MenuRepository) Get(id int) (models.Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetMenuByID

	menu := models.Menu{}

	row := repo.db.QueryRowContext(ctx, query, id)

	log.Println("GetMenuByID: ", query, id)

	err := row.Scan(
		&menu.ID,
		&menu.RestaurantId,
		&menu.Name,
		&menu.CreatedAt,
		&menu.UpdatedAt,
	)

	if err != nil {
		return menu, err
	}

	return menu, nil

}
