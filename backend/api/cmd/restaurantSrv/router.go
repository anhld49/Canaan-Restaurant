package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"backend/api/internal/handler/rest/public/middleware/authentication"
	"backend/api/internal/handler/rest/public/restaurant"
	"backend/api/internal/handler/rest/public/user"
)

type Router struct {
	userHandler       user.UserHandler
	restaurantHandler restaurant.RestaurantHandler
	router            *chi.Mux
}

// NewRouter: create new Router
func NewRouter(r user.UserHandler, res restaurant.RestaurantHandler) *Router {
	router := chi.NewRouter()
	return &Router{
		userHandler:       r,
		restaurantHandler: res,
		router:            router,
	}
}

func (r Router) adminRoutes() http.Handler {
	//Middleware with access rules for router.
	mux := chi.NewRouter()
	mux.Use(authentication.AuthRequired)
	mux.Get("/users", r.userHandler.List())
	mux.Post("/users", r.userHandler.Get())
	mux.Get("/users/{id}", r.userHandler.GetByID())

	// mux.Get("/restaurants", r.restaurantHandler.List())
	// mux.Get("/restaurants/{id}", r.restaurantHandler.Get())
	// mux.Put("/restaurants", r.restaurantHandler.Create())
	// mux.Patch("/restaurants/{id}", r.restaurantHandler.Update())
	// mux.Delete("/restaurants/{id}", r.restaurantHandler.Delete())

	return mux
}

// Routes: Router of users
func (r Router) routes() http.Handler {
	r.router.Use(middleware.Recoverer)
	r.router.Use(authentication.EnableCORS)

	r.router.Post("/authenticate", r.userHandler.Authenticate())
	r.router.Get("/refresh", r.userHandler.RefreshToken())
	r.router.Get("/logout", r.userHandler.Logout())

	r.router.Get("/restaurants", r.restaurantHandler.List())
	r.router.Get("/restaurants/{id}", r.restaurantHandler.Get())
	r.router.Put("/restaurants", r.restaurantHandler.Create())
	r.router.Patch("/restaurants/{id}", r.restaurantHandler.Update())
	r.router.Delete("/restaurants/{id}", r.restaurantHandler.Delete())

	// PROTECTED
	r.router.Mount("/admin", r.adminRoutes())

	return r.router
}
