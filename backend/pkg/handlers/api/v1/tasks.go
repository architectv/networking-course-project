package v1

import (
	"yak/backend/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (apiVX *ApiV1) registerTasksHandlers(router fiber.Router) {
	group := router.Group("/projects/:pid/boards/:bid/lists/:lid/tasks")
	group.Get("/", apiVX.getTasks)
	group.Post("/", apiVX.createTask)
	group.Get("/:tid", apiVX.getTask)
	group.Put("/:tid", apiVX.updateTask)
	group.Delete("/:tid", apiVX.deleteTask)
}

func (apiVX *ApiV1) getTasks(ctx *fiber.Ctx) error {
	implementMe()
	tasks := make([]models.Task, 0)
	return ctx.JSON(tasks)
}

func (apiVX *ApiV1) getTask(ctx *fiber.Ctx) error {
	implementMe()
	task := models.Task{}
	return ctx.JSON(task)
}

func (apiVX *ApiV1) createTask(ctx *fiber.Ctx) error {
	implementMe()
	task := models.Task{}
	return ctx.JSON(task)
}

func (apiVX *ApiV1) updateTask(ctx *fiber.Ctx) error {
	implementMe()
	task := models.Task{}
	return ctx.JSON(task)
}

func (apiVX *ApiV1) deleteTask(ctx *fiber.Ctx) error {
	implementMe()
	task := models.Task{}
	return ctx.JSON(task)
}
