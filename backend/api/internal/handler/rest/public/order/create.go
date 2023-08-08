package order

import (
	"backend/api/internal/handler/rest/public/middleware/authentication"
	"backend/api/internal/mod"
	"backend/api/internal/presenter"
	"backend/api/pkg/constants"
	"backend/api/pkg/utils"
	"log"
	"net/http"
	"time"
)

// Create: Create single Order
func (handler OrderHandler) Create() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := CreateOrderRequestPayload{}

		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		log.Println("requestPayload", requestPayload)

		userId := authentication.GetUserIdFromToken(r)
		orderDishes := []mod.OrderDish{}
		for _, value := range requestPayload.OrderDish {
			orderDishes = append(orderDishes, mod.OrderDish{
				DishId:   value.DishId,
				Quantity: value.Quantity,
			})
		}

		input := mod.Order{
			UserId:      userId,
			DriverId:    requestPayload.DriverId,
			OrderDishes: orderDishes,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		order, err := handler.controller.Create(input)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		orderDishesPres := []presenter.OrderDish{}
		for _, value := range order.OrderDishes {
			orderDishesPres = append(orderDishesPres, presenter.OrderDish{
				DishId:   value.DishId,
				Quantity: value.Quantity,
			})
		}

		resp := presenter.OrderResponse{
			Success: true,
			Message: constants.CREATE_ORDER_SUCCESSFULLY,
			Order: presenter.Order{
				ID:          order.ID,
				DriverId:    order.DriverId,
				Amount:      order.Amount,
				OrderDishes: orderDishesPres,
				CreatedAt:   order.CreatedAt,
				UpdatedAt:   order.UpdatedAt,
			},
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
