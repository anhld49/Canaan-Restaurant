package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"backend/api/internal/handler/rest/public/dish"
	"backend/api/internal/handler/rest/public/menu"
	"backend/api/internal/handler/rest/public/middleware/authentication"
	"backend/api/internal/handler/rest/public/order"
	"backend/api/internal/handler/rest/public/restaurant"
	"backend/api/internal/handler/rest/public/user"
)

type Router struct {
	userHandler       user.UserHandler
	restaurantHandler restaurant.RestaurantHandler
	menuHandler       menu.MenuHandler
	dishHandler       dish.DishHandler
	orderHandler      order.OrderHandler
	router            *chi.Mux
}

// NewRouter: create new Router
func NewRouter(r user.UserHandler, res restaurant.RestaurantHandler, me menu.MenuHandler, di dish.DishHandler, or order.OrderHandler) *Router {
	router := chi.NewRouter()
	return &Router{
		userHandler:       r,
		restaurantHandler: res,
		menuHandler:       me,
		dishHandler:       di,
		orderHandler:      or,
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

	mux.Get("/orders", r.orderHandler.List())
	mux.Get("/orders/{id}", r.orderHandler.Get())
	mux.Put("/orders", r.orderHandler.Create())
	mux.Delete("/orders/{id}", r.orderHandler.Delete())

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

	r.router.Get("/menus", r.menuHandler.List())
	r.router.Get("/menus/{id}", r.menuHandler.Get())
	r.router.Put("/menus", r.menuHandler.Create())
	r.router.Patch("/menus/{id}", r.menuHandler.Update())
	r.router.Delete("/menus/{id}", r.menuHandler.Delete())

	r.router.Get("/dishes", r.dishHandler.List())
	r.router.Get("/dishes/{id}", r.dishHandler.Get())
	r.router.Put("/dishes", r.dishHandler.Create())
	r.router.Patch("/dishes/{id}", r.dishHandler.Update())
	r.router.Delete("/dishes/{id}", r.dishHandler.Delete())

	// PROTECTED
	r.router.Mount("/admin", r.adminRoutes())

	return r.router
}
