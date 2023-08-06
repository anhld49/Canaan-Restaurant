package dish

import (
	"backend/api/internal/models"
	"backend/api/internal/presenter"
	"backend/api/pkg/constants"
	"backend/api/pkg/utils"
	"net/http"
	"time"
)

// Create: Create single Dish
func (handler DishHandler) Create() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := CreateDishRequestPayload{}
		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		input := models.Dish{
			MenuId:    requestPayload.MenuId,
			Name:      requestPayload.Name,
			Price:     requestPayload.Price,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		dish, err := handler.controller.Create(input)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		resp := presenter.DishResponse{
			Success: true,
			Message: constants.CREATE_MENU_SUCCESSFULLY,
			Dish: presenter.Dish{
				ID:        dish.ID,
				MenuId:    dish.MenuId,
				Name:      dish.Name,
				Price:     dish.Price,
				CreatedAt: dish.CreatedAt,
				UpdatedAt: dish.UpdatedAt,
			},
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
