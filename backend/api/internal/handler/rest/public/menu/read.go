package menu

import (
	"net/http"
	"strconv"

	"backend/api/internal/presenter"
	"backend/api/pkg/utils"

	"github.com/go-chi/chi/v5"
)

// List: get all menus
func (handler MenuHandler) List() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := handler.controller.List()
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}
		var menus []presenter.Menu
		for _, d := range data {
			menus = append(menus, presenter.Menu{
				ID:           d.ID,
				RestaurantId: d.RestaurantId,
				Name:         d.Name,
				CreatedAt:    d.CreatedAt,
				UpdatedAt:    d.UpdatedAt,
			})
		}

		utils.WriteJSON(w, http.StatusOK, menus)
	})
}

// Get: Get single menu by id
func (handler MenuHandler) Get() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		menu, err := handler.controller.Get(id)

		resp := presenter.Menu{
			ID:           menu.ID,
			RestaurantId: menu.RestaurantId,
			Name:         menu.Name,
			CreatedAt:    menu.CreatedAt,
			UpdatedAt:    menu.UpdatedAt,
		}
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
