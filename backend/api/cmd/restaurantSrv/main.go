package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	pkgErr "github.com/pkg/errors"

	"backend/api/internal/app/db"

	menuController "backend/api/internal/controller/menu"
	menuHandler "backend/api/internal/handler/rest/public/menu"
	menuRepository "backend/api/internal/repository/menu"

	restaurantController "backend/api/internal/controller/restaurant"
	restaurantHandler "backend/api/internal/handler/rest/public/restaurant"
	restaurantRepository "backend/api/internal/repository/restaurant"

	userController "backend/api/internal/controller/user"
	userHandler "backend/api/internal/handler/rest/public/user"
	userRepository "backend/api/internal/repository/user"
	"backend/api/pkg/constants"
)

func main() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5",
		os.Getenv(constants.DB_HOST), os.Getenv(constants.DB_PORT), os.Getenv(constants.DB_USER), os.Getenv(constants.DB_PASSWORD), os.Getenv(constants.DB_DATABASE))

	conn, err := db.Connect(dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	initServer(conn)
}

func initServer(conn *sql.DB) {
	userRepo := userRepository.NewUserRepo(conn)
	userController := userController.NewUserController(userRepo)
	userHandler := userHandler.NewUserHandler(*userController)

	restaurantRepo := restaurantRepository.NewRestaurantRepo(conn)
	restaurantController := restaurantController.NewRestaurantController(*restaurantRepo)
	restaurantHandler := restaurantHandler.NewRestaurantHandler(*restaurantController)

	menuRepo := menuRepository.NewMenuRepo(conn)
	menuController := menuController.NewMenuController(*menuRepo)
	menuHandler := menuHandler.NewMenuHandler(*menuController)

	router := NewRouter(*userHandler, *restaurantHandler, *menuHandler)

	log.Println("Starting application on port", os.Getenv(constants.API_PORT))

	// start a web server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv(constants.API_PORT)), router.routes()); err != nil {
		log.Fatal(pkgErr.WithStack(err))
	}
}
