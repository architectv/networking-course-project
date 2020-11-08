package v1

import (
	"yak/backend/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (apiVX *ApiV1) registerListsHandlers(router fiber.Router) {
	group := router.Group("/projects/:pid/boards/:bid/lists")
	group.Get("/", apiVX.getLists)
	group.Post("/", apiVX.createList)
	group.Get("/:lid", apiVX.getList)
	group.Put("/:lid", apiVX.updateList)
	group.Delete("/:lid", apiVX.deleteList)
}

func (apiVX *ApiV1) getLists(ctx *fiber.Ctx) error {
	implementMe()
	lists := make([]models.TaskList, 0)
	return ctx.JSON(lists)
}

func (apiVX *ApiV1) getList(ctx *fiber.Ctx) error {
	implementMe()
	list := models.TaskList{}
	return ctx.JSON(list)
}

func (apiVX *ApiV1) createList(ctx *fiber.Ctx) error {
	implementMe()
	list := models.TaskList{}
	return ctx.JSON(list)
}

func (apiVX *ApiV1) updateList(ctx *fiber.Ctx) error {
	implementMe()
	list := models.TaskList{}
	return ctx.JSON(list)
}

func (apiVX *ApiV1) deleteList(ctx *fiber.Ctx) error {
	implementMe()
	list := models.TaskList{}
	return ctx.JSON(list)
}
