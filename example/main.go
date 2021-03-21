package main

import (
	"github.com/gofiber/fiber/v2"
	redoc "github.com/natebwangsut/fiber-redoc"
	_ "github.com/natebwangsut/fiber-redoc/example/docs" // This is required as `swag init` would generate docs package where swag interface will be exported
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})

	app.Get("/docs/*", redoc.Handler)
	app.Listen(":8080")
}
