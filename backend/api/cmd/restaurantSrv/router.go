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

func (r UserRouter) adminRoutes() http.Handler {
	//Middleware with access rules for router.
	mux := chi.NewRouter()
	mux.Use(authentication.AuthRequired)
	mux.Get("/users", r.handler.List())

	return mux
}

// Routes: Router of users
func (r UserRouter) routes() http.Handler {
	r.router.Use(middleware.Recoverer)
	r.router.Use(authentication.EnableCORS)

	r.router.Post("/authenticate", r.handler.Authenticate())
	r.router.Get("/refresh", r.handler.RefreshToken())
	r.router.Get("/logout", r.handler.Logout())

	// PROTECTED
	r.router.Mount("/admin", r.adminRoutes())

	return r.router
}
