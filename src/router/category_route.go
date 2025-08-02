package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(v1 fiber.Router, p service.CategoryService) {
	categoryController := controller.NewCategoryController(p)

	category := v1.Group("/category")

	category.Get("list", categoryController.GetCategoryList)
	category.Get(":id", categoryController.GetCategoryByCategoryID)
	category.Post("", categoryController.CreateCategory)

}
