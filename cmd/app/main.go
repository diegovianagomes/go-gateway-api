package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/diegovianagomes/go-gateway-api/internal/repository"
	"github.com/diegovianagomes/go-gateway-api/internal/service"
	"github.com/diegovianagomes/go-gateway-api/internal/web/server"
	"github.com/joho/godotenv"
)

func getEnv(key, defaultValue string) string{
	if value := os.Getenv(key); value != ""{
		return value
	}
	return defaultValue
}



func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Set up connection to PostgreSQL using ambient variables
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	// Initialize application layers (repository -> service -> server)
	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	// Configure and start the HTTP server
	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(accountService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}


