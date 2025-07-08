package config

import "fmt"

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
    DB_USER     = "postgres"	
    DB_PASSWORD = "123"
	DB_NAME     = "todo_db"
)

func GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT,
	)
}