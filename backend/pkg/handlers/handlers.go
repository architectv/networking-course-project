package handlers

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/handlers/api"

func RegisterHandlers(router fiber.Router) {
	api.RegisterHandlers(router)
}
