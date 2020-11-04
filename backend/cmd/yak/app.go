package yak

import (
	"yak/backend/pkg/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateApp() {
	app := fiber.New()
	app.Use(logger.New())
	handlers.RegisterHandlers(app)
	app.Listen(":8001")
}
