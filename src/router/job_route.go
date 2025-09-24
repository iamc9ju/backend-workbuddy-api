package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func JobRoutes(v1 fiber.Router, p service.JobService) {
	projectController := controller.NewProjectController(p)

	project := v1.Group("/project")

	project.Get("list", projectController.GetProjectList)
	project.Get(":slug", projectController.GetProjectBySlug)
	project.Post("", projectController.CreateProject)
	project.Get("owner/:id", projectController.GetProjectsByOwnerID)

}
