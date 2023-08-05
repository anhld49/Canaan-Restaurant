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
	controller "backend/api/internal/controller/user"
	handler "backend/api/internal/handler/rest/public/user"
	repository "backend/api/internal/repository/user"
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
	userRepo := repository.NewUserRepo(conn)
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(*userController)
	userRouter := NewUserRouter(*userHandler)

	log.Println("Starting application on port", os.Getenv(constants.API_PORT))

	// start a web server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv(constants.API_PORT)), userRouter.routes()); err != nil {
		log.Fatal(pkgErr.WithStack(err))
	}
}
