package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"

	"backend/api/internal/handler/rest/public/middleware/authentication"
	"backend/api/internal/presenter"
	"backend/api/pkg/constants"
	"backend/api/pkg/utils"
)

// Authenticate: Authenticate a user
func (handler UserHandler) Authenticate() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := AuthenticationRequestPayload{}
		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		// validate user against database
		user, err := handler.controller.Get(requestPayload.Email)
		if err != nil {
			utils.ErrorJSON(w, errors.New(constants.INVALID_CREDENTIAL), http.StatusBadRequest)
			return
		}

		// check password
		valid, err := user.PasswordMatches(requestPayload.Password)
		if err != nil || !valid {
			utils.ErrorJSON(w, errors.New(constants.INVALID_CREDENTIAL), http.StatusBadRequest)
			return
		}

		// create a jwt user
		u := authentication.JwtUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		// generate tokens
		tokens, err := authentication.AUTH_CONFIG.GenerateTokenPair(&u)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		refreshCookie := authentication.AUTH_CONFIG.GetRefreshCookie(tokens.RefreshToken)
		http.SetCookie(w, refreshCookie)

		utils.WriteJSON(w, http.StatusAccepted, tokens)
	})
}

// RefreshToken: Generate a RefreshToken
func (handler UserHandler) RefreshToken() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, cookie := range r.Cookies() {
			if cookie.Name == authentication.AUTH_CONFIG.CookieName {
				claims := &authentication.Claims{}
				refreshToken := cookie.Value

				// parse the token to get the claims
				_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(authentication.AUTH_CONFIG.Secret), nil
				})
				if err != nil {
					utils.ErrorJSON(w, errors.New(constants.UNAUTHORIZED), http.StatusUnauthorized)
					return
				}

				// get the user id from the token claims
				userID, err := strconv.Atoi(claims.Subject)
				if err != nil {
					utils.ErrorJSON(w, errors.New(constants.UNKNOWN_USER), http.StatusUnauthorized)
					return
				}

				user, err := handler.controller.GetByID(userID)
				if err != nil {
					utils.ErrorJSON(w, errors.New(constants.UNKNOWN_USER), http.StatusUnauthorized)
					return
				}

				u := authentication.JwtUser{
					ID:        user.ID,
					FirstName: user.FirstName,
					LastName:  user.LastName,
				}

				tokenPairs, err := authentication.AUTH_CONFIG.GenerateTokenPair(&u)
				if err != nil {
					utils.ErrorJSON(w, errors.New(constants.GENERATING_TOKEN), http.StatusUnauthorized)
					return
				}

				http.SetCookie(w, authentication.AUTH_CONFIG.GetRefreshCookie(tokenPairs.RefreshToken))

				utils.WriteJSON(w, http.StatusOK, tokenPairs)

			}
		}
	})
}

// Logout: Logout after authentication
func (handler UserHandler) Logout() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, authentication.AUTH_CONFIG.GetExpiredRefreshCookie())
		w.WriteHeader(http.StatusAccepted)
	})
}

// List: get all users
func (handler UserHandler) List() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := handler.controller.List()
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}
		var users []presenter.User
		for _, d := range data {
			users = append(users, presenter.User{
				ID:        d.ID,
				FirstName: d.FirstName,
				LastName:  d.LastName,
				Email:     d.Email,
				Role:      d.Role,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}

		utils.WriteJSON(w, http.StatusOK, users)
	})
}

// Get: Get single user by email
func (handler UserHandler) Get() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := EmailRequestPayload{}
		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		user, err := handler.controller.Get(requestPayload.Email)

		resp := presenter.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}

// Get: Get single user by id
func (handler UserHandler) GetByID() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		user, err := handler.controller.GetByID(id)

		resp := presenter.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
