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

type ProjectController struct {
	project_service service.ProjectService
}

func NewProjectController(project_service service.ProjectService) *ProjectController {
	return &ProjectController{project_service: project_service}
}

func (ctl *ProjectController) GetProjectList(c *fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancle()
	projects, err := ctl.project_service.GetProjectList(ctx)

	if err != nil {
		utils.Log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get projects",
			Errors:  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithProjectList{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Get all projects successfully",
		Project: projects,
	})
}

func (ctl *ProjectController) CreateProject(c *fiber.Ctx) error {
	var body model.ProjectCreate
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorDetails{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Get owner ID
	//  ownerID := getUserIDFromContext(c)
	project, err := ctl.project_service.CreateProject(c.Context(), body, body.OwnerID)
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
			Message: "Failed to create project",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(response.SuccessWithProject{
		Code:    fiber.StatusCreated,
		Status:  "success",
		Message: "Project created successfully",
		Project: *project,
	})
}

func (ctl *ProjectController) GetProjectByProjectID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// ดึง ID จาก URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Common{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "Invalid project ID",
		})
	}

	project, err := ctl.project_service.GetProjectByProjectID(ctx, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(response.ErrorDetails{
				Code:    fiber.StatusNotFound,
				Status:  "error",
				Message: "Project not found",
				Errors:  err,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get project",
			Errors:  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithProject{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Project retrieved successfully",
		Project: *project,
	})
}

func (ctl *ProjectController) GetProjectBySlug(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// ดึง slug จาก URL parameter
	slug := c.Params("slug")
	project, err := ctl.project_service.GetProjectBySlug(ctx, slug)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(response.ErrorDetails{
				Code:    fiber.StatusNotFound,
				Status:  "error",
				Message: "Project not found",
				Errors:  err,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get project",
			Errors:  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithProject{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Project retrieved successfully",
		Project: *project,
	})
}

func (c *ProjectController) GetProjectsByOwnerID(ctx *fiber.Ctx) error {
	// ownerID := getUserIDFromContext(ctx) // ดึงจาก JWT หรือ session
	id, err := ctx.ParamsInt("id")

	projects, err := c.project_service.GetProjectsByOwnerID(ctx.Context(), uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no projects found") {
			return ctx.Status(fiber.StatusOK).JSON(response.ErrorDetails{
				Code:    fiber.StatusNotFound,
				Status:  "error",
				Message: "Project not found",
				Errors:  err,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get project",
			Errors:  err,
		})
	}

	// return ctx.JSON(projects)
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessWithProjectList{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Project retrieved successfully",
		Project: projects,
	})
}
