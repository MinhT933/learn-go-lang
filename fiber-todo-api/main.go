package main

import (
	"fiber-todo-api/database"
	"fiber-todo-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()
	routes.SetupTaskRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Todo API đang chạy!")
	})

	// ❗ Dòng quan trọng nhất để giữ server chạy
	log.Fatal(app.Listen(":8081"))
}
