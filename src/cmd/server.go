package main

import (
	"log"

	"github.com/alaa-aqeel/govalid/src/api/routes"
	"github.com/alaa-aqeel/govalid/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Pkg
	repository.InitializeDatabase()

	// Initialize Gin
	r := gin.Default()
	routes.Routes(r)
	r.Run(":8080")
}
