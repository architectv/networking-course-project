package v1

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/models"

func registerTasksHandlers(router fiber.Router) {
	group := router.Group("/projects/:pid/boards/:bid/lists/:lid/tasks")
	group.Get("/", getTasks)
	group.Post("/", createTask)
	group.Get("/:tid", getTask)
	group.Put("/:tid", updateTask)
	group.Delete("/:tid", deleteTask)
}

func getTasks(ctx *fiber.Ctx) error {
	implementMe()
	tasks := make([]models.Task, 0)
	return ctx.JSON(tasks)
}

func getTask(ctx *fiber.Ctx) error {
	implementMe()
	task := models.Task{}
	return ctx.JSON(task)
}

func createTask(ctx *fiber.Ctx) error {
	implementMe()
	task := models.Task{}
	return ctx.JSON(task)
}

func updateTask(ctx *fiber.Ctx) error {
	implementMe()
	task := models.Task{}
	return ctx.JSON(task)
}

func deleteTask(ctx *fiber.Ctx) error {
	implementMe()
	task := models.Task{}
	return ctx.JSON(task)
}
