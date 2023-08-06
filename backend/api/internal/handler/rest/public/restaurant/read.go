package restaurant

import (
	"net/http"
	"strconv"

	"backend/api/internal/presenter"
	"backend/api/pkg/utils"

	"github.com/go-chi/chi/v5"
)

// List: get all restaurants
func (handler RestaurantHandler) List() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := handler.controller.List()
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}
		var restaurants []presenter.Restaurant
		for _, d := range data {
			restaurants = append(restaurants, presenter.Restaurant{
				ID:        d.ID,
				OwnerId:   d.OwnerId,
				Name:      d.Name,
				Address:   d.Address,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}

		utils.WriteJSON(w, http.StatusOK, restaurants)
	})
}

// Get: Get single restaurant by id
func (handler RestaurantHandler) Get() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		restaurant, err := handler.controller.Get(id)

		resp := presenter.Restaurant{
			ID:        restaurant.ID,
			OwnerId:   restaurant.OwnerId,
			Name:      restaurant.Name,
			Address:   restaurant.Address,
			CreatedAt: restaurant.CreatedAt,
			UpdatedAt: restaurant.UpdatedAt,
		}
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
