package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"backend/api/internal/handler/rest/public/middleware/authentication"
	"backend/api/internal/handler/rest/public/user"
)

// UserRouter: User Router
type UserRouter struct {
	handler user.UserHandler
	router  *chi.Mux
}

// NewUserRouter: create new user Router
func NewUserRouter(r user.UserHandler) *UserRouter {
	router := chi.NewRouter()
	return &UserRouter{
		handler: r,
		router:  router,
	}
}

// Routes: Router of users
func (r UserRouter) routes() http.Handler {
	r.router.Use(middleware.Recoverer)
	r.router.Use(authentication.EnableCORS)

	r.router.Post("/authenticate", r.handler.Authenticate())
	r.router.Get("/refresh", r.handler.RefreshToken())
	r.router.Get("/logout", r.handler.Logout())

	r.router.Route("/owner", func(mux chi.Router) {
		mux.Use(authentication.AuthRequired)

		// r.router.Get("/menus", app.MovieCatalog)
	})

	r.router.Get("/users", r.handler.List())
	r.router.Post("/users", r.handler.Get())
	r.router.Post("/users/id", r.handler.GetByID())
	// r.router.Post("/invite", r.handler.CreateFriendship())
	// r.router.Post("/subscribe", r.handler.CreateSubscribe())
	// r.router.Post("/blocks", r.handler.CreateBlock())

	return r.router
}
