package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func JobRoutes(v1 fiber.Router, p service.JobService) {
	jobController := controller.NewJobController(p)

	job := v1.Group("/job")

	job.Get("list", jobController.GetJobList)
	job.Get(":slug", jobController.GetJobBySlug)
	job.Post("", jobController.CreateJob)
	job.Get("owner/:id", jobController.GetJobByOwnerID)

}
