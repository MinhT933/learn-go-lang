package routes

import (
	"fiber-todo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(app *fiber.App) {
	// task := app.Group("/todos")
	task:= app.Group("/tasks")
	task.Get("/", handlers.GetAllTasks)
	task.Post("/", handlers.CreateTask)
	task.Get("/:id", handlers.GetTaskByID)
	task.Delete("/:id", handlers.DeleteTaskByID)
	task.Put("/:id", handlers.UpdateTaskByID)
}