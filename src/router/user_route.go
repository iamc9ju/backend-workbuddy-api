package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(v1 fiber.Router, u service.UserService) {
	userController := controller.NewUserController(u)

	user := v1.Group("/user")

	user.Get("/list", userController.GetAllUsers)
	// user.Post("/", m.Auth(u, "manageUsers"), userController.CreateUser)
	// user.Get("/:userId", m.Auth(u, "getUsers"), userController.GetUserByID)
	// user.Patch("/:userId", m.Auth(u, "manageUsers"), userController.UpdateUser)
	// user.Delete("/:userId", m.Auth(u, "manageUsers"), userController.DeleteUser)
}
