package handlers

import (
	"yak/backend/pkg/handlers/api"
	"yak/backend/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) RegisterHandlers(router fiber.Router) {
	api := api.NewApi(h.services)
	api.RegisterHandlers(router)
}
