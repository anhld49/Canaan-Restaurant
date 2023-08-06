package restaurant

import (
	"backend/api/internal/models"
	"backend/api/internal/presenter"
	"backend/api/pkg/constants"
	"backend/api/pkg/utils"
	"net/http"
	"time"
)

// Create: Create single restaurant
func (handler RestaurantHandler) Create() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := CreateRestaurantRequestPayload{}
		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		input := models.Restaurant{
			OwnerId:   requestPayload.OwnerId,
			Name:      requestPayload.Name,
			Address:   requestPayload.Address,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		restaurant, err := handler.controller.Create(input)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		resp := presenter.RestaurantResponse{
			Success: true,
			Message: constants.CREATE_RESTAURANT_SUCCESSFULLY,
			Restaurant: presenter.Restaurant{
				ID:        restaurant.ID,
				OwnerId:   restaurant.OwnerId,
				Name:      restaurant.Name,
				Address:   restaurant.Address,
				CreatedAt: restaurant.CreatedAt,
				UpdatedAt: restaurant.UpdatedAt,
			},
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
