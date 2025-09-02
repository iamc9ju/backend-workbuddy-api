package controller

import (
	"app/src/response"
	"app/src/service"
	"app/src/utils"

	"github.com/gofiber/fiber/v2"
)

type LanguageController struct {
	service service.LanguageService
}

func NewLanguageController(service service.LanguageService) *LanguageController {
	return &LanguageController{service: service}
}

func (lc *LanguageController) GetAllLanguages(c *fiber.Ctx) error {
	languages, err := lc.service.ListAllLanguages()

	if err != nil {
		utils.Log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get languages",
			Errors:  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithLanguageList{
		Code:      fiber.StatusOK,
		Status:    "success",
		Message:   "success to get languages",
		Languages: languages,
	})
}
