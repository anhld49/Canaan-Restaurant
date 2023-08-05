package db

import (
	"fmt"
	"testing"
)

var (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_DATABASE = "backend"
	DB_PORT     = "5432"
	dsn         = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

func TestConnect(t *testing.T) {
	dsn := fmt.Sprintf(dsn, DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE)
	_, err := openDB(dsn)

	if err != nil {
		t.Errorf("Db Connect reports error, expected nil, but got %s", err.Error())
	}
}
