package order

import (
	"backend/api/internal/presenter"
	"backend/api/pkg/constants"
	"backend/api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Delete: Delete single order
func (handler OrderHandler) Delete() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		err = handler.controller.Delete(id)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		resp := presenter.DeleteOrderResponse{
			Success: true,
			Message: constants.DELETE_ORDER_SUCCESSFULLY,
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
