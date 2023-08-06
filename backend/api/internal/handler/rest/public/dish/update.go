package dish

import (
	"backend/api/internal/models"
	"backend/api/internal/presenter"
	"backend/api/pkg/constants"
	"backend/api/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// Update: Update single dish
func (handler DishHandler) Update() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		requestPayload := UpdateDishRequestPayload{}
		err = utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		input := models.Dish{
			ID:        id,
			MenuId:    requestPayload.MenuId,
			Name:      requestPayload.Name,
			Price:     requestPayload.Price,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		dish, err := handler.controller.Update(input)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		resp := presenter.DishResponse{
			Success: true,
			Message: constants.UPDATE_MENU_SUCCESSFULLY,
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
