package yak

import "github.com/gofiber/fiber/v2"

func CreateApp() {
	api := fiber.New()
	api.Listen(":8001")
}
