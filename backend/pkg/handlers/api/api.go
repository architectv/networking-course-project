package api

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/handlers/api/v1"

func RegisterHandlers(group fiber.Router) {
	api_v1 := group.Group("/v1")
	v1.RegisterHandlers(api_v1)
}
