package controller

import (
	"app/src/model"
	"app/src/response"
	"app/src/service"
	"app/src/utils"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	category_service service.CategoryService
}

func NewCategoryController(category_service service.CategoryService) *CategoryController {
	return &CategoryController{category_service: category_service}
}

func (ctl *CategoryController) GetCategoryList(c *fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancle()
	category, err := ctl.category_service.GetCategoryList(ctx)

	if err != nil {
		utils.Log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get category",
			Errors:  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithCategory{
		Code:     fiber.StatusOK,
		Status:   "success",
		Message:  "Get all category successfully",
		Category: category,
	})
}

func (ctl *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var body model.CategoryCreate
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorDetails{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	category, err := ctl.category_service.CreateCategory(c.Context(), body)
	if err != nil {
		if errors.Is(err, validator.ValidationErrors{}) {
			return c.Status(fiber.StatusBadRequest).JSON(response.ErrorDetails{
				Code:    fiber.StatusBadRequest,
				Status:  "error",
				Message: "Validation error",
				Errors:  err.(validator.ValidationErrors),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.Common{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to create category",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(response.SuccessWithCategory{
		Code:     fiber.StatusCreated,
		Status:   "success",
		Message:  "Category created successfully",
		Category: []model.Category{*category},
	})
}

func (ctl *CategoryController) GetCategoryByCategoryID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// ดึง ID จาก URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Common{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "Invalid category ID",
		})
	}

	category, err := ctl.category_service.GetCategoryByCategoryID(ctx, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(response.ErrorDetails{
				Code:    fiber.StatusNotFound,
				Status:  "error",
				Message: "Category not found",
				Errors:  err,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get category",
			Errors:  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithCategory{
		Code:     fiber.StatusOK,
		Status:   "success",
		Message:  "Category retrieved successfully",
		Category: []model.Category{*category},
	})
}
