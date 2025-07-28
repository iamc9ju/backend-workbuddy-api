package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(v1 fiber.Router, p service.ProjectService) {
	projectController := controller.NewProjectController(p)

	project := v1.Group("/project")

	project.Get("/getProjectList", projectController.GetProjectList)
	project.Get("/getProjectByProjectId/:id", projectController.GetProjectByProjectID)
	project.Post("/createProject", projectController.CreateProject)
	project.Get("/getProjectByOwnerId/:id", projectController.GetProjectsByOwnerID)

}
