package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func LanguageRoutes(v1 fiber.Router, l service.LanguageService) {
	languageController := controller.NewLanguageController(l)

	language := v1.Group("/languages")
	language.Get("list", languageController.GetAllLanguages)

}
