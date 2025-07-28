package router

import (
	"app/src/config"
	"app/src/repository"
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {
	validate := validation.Validator()

	// healthCheckService := service.NewHealthCheckService(db)
	// emailService := service.NewEmailService()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, validate)
	// tokenService := service.NewTokenService(db, validate, userService)
	// authService := service.NewAuthService(db, validate, userService, tokenService)

	v1 := app.Group("/v1")

	// HealthCheckRoutes(v1, healthCheckService)
	// AuthRoutes(v1, authService, userService, tokenService, emailService)
	UserRoutes(v1, userService)
	// TODO: add another routes here...

	// Development-only routes
	if !config.IsProd {
		DocsRoutes(v1)
	}
}
