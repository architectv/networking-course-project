package handlers

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/handlers/api"

func RegisterHandlers(group fiber.Router) {
	api_v1 := group.Group("/v1")
	api.RegisterHandlers(api_v1)
}
