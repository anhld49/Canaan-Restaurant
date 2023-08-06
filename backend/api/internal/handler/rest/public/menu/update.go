package menu

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

// Update: Update single menu
func (handler MenuHandler) Update() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		requestPayload := UpdateMenuRequestPayload{}
		err = utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		input := models.Menu{
			ID:           id,
			RestaurantId: requestPayload.RestaurantId,
			Name:         requestPayload.Name,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		menu, err := handler.controller.Update(input)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		resp := presenter.MenuResponse{
			Success: true,
			Message: constants.UPDATE_MENU_SUCCESSFULLY,
			Menu: presenter.Menu{
				ID:           menu.ID,
				RestaurantId: menu.RestaurantId,
				Name:         menu.Name,
				CreatedAt:    menu.CreatedAt,
				UpdatedAt:    menu.UpdatedAt,
			},
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
