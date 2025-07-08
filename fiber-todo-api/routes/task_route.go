package routes

import (
	"fiber-todo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(app *fiber.App) {
	task := app.Group("/todos")
	task.Get("/", handlers.GetAllTasks)
	task.Post("/", handlers.CreateTask)
}