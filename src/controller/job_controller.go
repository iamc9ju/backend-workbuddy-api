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
	project_service service.JobService
}

func NewJobController(project_service service.JobService) *ProjectController {
	return &ProjectController{project_service: project_service}
}

func (ctl *ProjectController) GetJobList(c *fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancle()
	jobs, err := ctl.project_service.GetJobList(ctx)

	if err != nil {
		utils.Log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get jobs",
			Errors:  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithJobList{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Get all jobs successfully",
		Job:     jobs,
	})
}

func (ctl *ProjectController) CreateJob(c *fiber.Ctx) error {
	var body model.JobCreate
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorDetails{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Get owner ID
	//  ownerID := getUserIDFromContext(c)
	job, err := ctl.project_service.CreateJob(c.Context(), body, body.OwnerID)
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
	return c.Status(fiber.StatusCreated).JSON(response.SuccessWithJob{
		Code:    fiber.StatusCreated,
		Status:  "success",
		Message: "Job created successfully",
		Job:     *job,
	})
}

func (ctl *ProjectController) GetJobByJobID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// ดึง ID จาก URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Common{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "Invalid job ID",
		})
	}

	job, err := ctl.project_service.GetJobByJobID(ctx, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(response.ErrorDetails{
				Code:    fiber.StatusNotFound,
				Status:  "error",
				Message: "Job not found",
				Errors:  err,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get job",
			Errors:  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithJob{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Job retrieved successfully",
		Job:     *job,
	})
}

func (ctl *ProjectController) GetJobBySlug(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// ดึง slug จาก URL parameter
	slug := c.Params("slug")
	job, err := ctl.project_service.GetJobBySlug(ctx, slug)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(response.ErrorDetails{
				Code:    fiber.StatusNotFound,
				Status:  "error",
				Message: "Job not found",
				Errors:  err,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get job",
			Errors:  err,
		})

	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithJob{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Job retrieved successfully",
		Job:     *job,
	})
}

func (c *ProjectController) GetJobByOwnerID(ctx *fiber.Ctx) error {
	// ownerID := getUserIDFromContext(ctx) // ดึงจาก JWT หรือ session
	id, err := ctx.ParamsInt("id")

	jobs, err := c.project_service.GetJobByOwnerID(ctx.Context(), uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no jobs found") {
			return ctx.Status(fiber.StatusOK).JSON(response.ErrorDetails{
				Code:    fiber.StatusNotFound,
				Status:  "error",
				Message: "Job not found",
				Errors:  err,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorDetails{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to get job",
			Errors:  err,
		})
	}

	// return ctx.JSON(projects)
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessWithJobList{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Job retrieved successfully",
		Job:     jobs,
	})
}
