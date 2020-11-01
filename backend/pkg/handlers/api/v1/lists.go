package v1

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/models"

func registerListsHandlers(router fiber.Router) {
	group := router.Group("/projects/:pid/boards/:bid/lists")
	group.Get("/", getLists)
	group.Post("/", createList)
	group.Get("/:lid", getList)
	group.Put("/:lid", updateList)
	group.Delete("/:lid", deleteList)
}

func getLists(ctx *fiber.Ctx) error {
	implementMe()
	lists := make([]models.TaskList, 0)
	return ctx.JSON(lists)
}

func getList(ctx *fiber.Ctx) error {
	implementMe()
	list := models.TaskList{}
	return ctx.JSON(list)
}

func createList(ctx *fiber.Ctx) error {
	implementMe()
	list := models.TaskList{}
	return ctx.JSON(list)
}

func updateList(ctx *fiber.Ctx) error {
	implementMe()
	list := models.TaskList{}
	return ctx.JSON(list)
}

func deleteList(ctx *fiber.Ctx) error {
	implementMe()
	list := models.TaskList{}
	return ctx.JSON(list)
}
