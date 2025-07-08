package main

// @title Fiber Todo API
// @version 1.0
// @description REST API sử dụng Fiber và Swagger
// @host localhost:3000
// @BasePath /

import (
	"fiber-todo-api/database"
	"fiber-todo-api/routes"

	_ "fiber-todo-api/docs" // 👈 Quan trọng: import docs cho Swagger

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // ✅ Fiber's logger
	fiberSwagger "github.com/swaggo/fiber-swagger"  // 👈 chứa hàm Handler
)
func main() {
	app := fiber.New()
    app.Use(logger.New())
	database.ConnectDB()
	routes.SetupTaskRoutes(app)


    app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Listen(":3000")
}
