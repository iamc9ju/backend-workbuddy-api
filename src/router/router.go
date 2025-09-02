package router

import (
	"app/src/config"
	"app/src/repository"
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	validate := validation.Validator()

	// healthCheckService := service.NewHealthCheckService(db)
	// emailService := service.NewEmailService()
	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	languagesRepo := repository.NewProgrammingLanguageRepository(db)
	userService := service.NewUserService(userRepo, validate)
	projectService := service.NewProjectService(projectRepo, validate)
	categoryService := service.NewCategoryService(categoryRepo, validate)
	// tokenService := service.NewTokenService(db, validate, userService)
	// authService := service.NewAuthService(db, validate, userService, tokenService)
	languageService := service.NewLanguageService(languagesRepo, validate)

	v1 := app.Group("/v1")

	// HealthCheckRoutes(v1, healthCheckService)
	// AuthRoutes(v1, authService, userService, tokenService, emailService)
	UserRoutes(v1, userService)
	ProjectRoutes(v1, projectService)
	CategoryRoutes(v1, categoryService)
	LanguageRoutes(v1, languageService)
	// TODO: add another routes here...

	// Development-only routes
	if !config.IsProd {
		DocsRoutes(v1)
	}
}
