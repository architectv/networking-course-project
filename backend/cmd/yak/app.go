package yak

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/handlers"

func CreateApp() {
	app := fiber.New()
	handlers.RegisterHandlers(app)
	app.Listen(":8001")
}
