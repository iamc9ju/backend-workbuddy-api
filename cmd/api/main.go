package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iamc9ju/backend-workbuddy-api/internal/adapters/input/http/router"
	"github.com/iamc9ju/backend-workbuddy-api/internal/config"
)

func main() {
	//Initialize Gin router
	err := godotenv.Load()
	if err != nil {
        log.Fatal("Error loading .env file")
    }

	route := gin.Default()

	if os.Getenv("STATE") == "DEV" {
		route.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Typically include more methods
			AllowHeaders:     []string{"*"},                                       // Allow important headers
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	} else if os.Getenv("STATE") == "STG" {
		route.Use(cors.New(cors.Config{
			AllowOrigins:     []string{""},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	} else if os.Getenv("STATE") == "PROD" {
		route.Use(cors.New(cors.Config{
			AllowOrigins:     []string{""},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	// AWS configuration
	// awsconfig.LoadAWSConfig()

	// Allow CORS
	corsConfig := config.AllowCor()
	handler := corsConfig.Handler(route)

	// Configure routes
	router.Router(route)

	// Get port from .env variable
	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatalf("Port should be set...")
	}
	log.Printf("Starting server on port %s", PORT)

	// Run application
	err = http.ListenAndServe(PORT, handler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
