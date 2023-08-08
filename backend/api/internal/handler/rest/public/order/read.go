package order

import (
	"net/http"
	"strconv"

	"backend/api/internal/handler/rest/public/middleware/authentication"
	"backend/api/internal/presenter"
	"backend/api/pkg/utils"

	"github.com/go-chi/chi/v5"
)

// List: get all orders
func (handler OrderHandler) List() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := authentication.GetUserIdFromToken(r)

		data, err := handler.controller.List(userId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		var orders []presenter.Order
		for _, order := range data {

			orderDishesPres := []presenter.OrderDish{}
			for _, value := range order.OrderDishes {
				orderDishesPres = append(orderDishesPres, presenter.OrderDish{
					DishId:   value.DishId,
					Quantity: value.Quantity,
				})
			}
			orders = append(orders, presenter.Order{
				ID:          order.ID,
				DriverId:    order.DriverId,
				Amount:      order.Amount,
				OrderDishes: orderDishesPres,
				CreatedAt:   order.CreatedAt,
				UpdatedAt:   order.UpdatedAt,
			})
		}

		utils.WriteJSON(w, http.StatusOK, orders)
	})
}

// Get: Get single order by id
func (handler OrderHandler) Get() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		order, err := handler.controller.Get(id)

		orderDishesPres := []presenter.OrderDish{}
		for _, value := range order.OrderDishes {
			orderDishesPres = append(orderDishesPres, presenter.OrderDish{
				DishId:   value.DishId,
				Quantity: value.Quantity,
			})
		}

		resp := presenter.Order{
			ID:          order.ID,
			DriverId:    order.DriverId,
			Amount:      order.Amount,
			OrderDishes: orderDishesPres,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		}
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
