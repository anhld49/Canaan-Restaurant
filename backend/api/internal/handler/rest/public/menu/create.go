package menu

import (
	"backend/api/internal/models"
	"backend/api/internal/presenter"
	"backend/api/pkg/constants"
	"backend/api/pkg/utils"
	"net/http"
	"time"
)

// Create: Create single Menu
func (handler MenuHandler) Create() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := CreateMenuRequestPayload{}
		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		input := models.Menu{
			RestaurantId: requestPayload.RestaurantId,
			Name:         requestPayload.Name,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		menu, err := handler.controller.Create(input)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		resp := presenter.MenuResponse{
			Success: true,
			Message: constants.CREATE_MENU_SUCCESSFULLY,
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
