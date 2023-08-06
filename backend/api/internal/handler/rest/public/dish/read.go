package dish

import (
	"net/http"
	"strconv"

	"backend/api/internal/presenter"
	"backend/api/pkg/utils"

	"github.com/go-chi/chi/v5"
)

// List: get all dishs
func (handler DishHandler) List() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := handler.controller.List()
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}
		var dishs []presenter.Dish
		for _, d := range data {
			dishs = append(dishs, presenter.Dish{
				ID:        d.ID,
				MenuId:    d.MenuId,
				Name:      d.Name,
				Price:     d.Price,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}

		utils.WriteJSON(w, http.StatusOK, dishs)
	})
}

// Get: Get single dish by id
func (handler DishHandler) Get() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		dish, err := handler.controller.Get(id)

		resp := presenter.Dish{
			ID:        dish.ID,
			MenuId:    dish.MenuId,
			Name:      dish.Name,
			Price:     dish.Price,
			CreatedAt: dish.CreatedAt,
			UpdatedAt: dish.UpdatedAt,
		}
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
