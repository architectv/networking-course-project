package api

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/handlers/api/v1"

func RegisterHandlers(router fiber.Router) {
	api := router.Group("/api")
	v1.RegisterHandlers(api)
}
