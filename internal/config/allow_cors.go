package config

import (
	"net/http"

	"github.com/rs/cors"
)

func AllowCor() *cors.Cors {
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "request-id", "application-channel", "transaction-dateTime"},
		AllowCredentials: true,
	})

	return corsConfig
}
