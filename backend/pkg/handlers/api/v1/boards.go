package v1

import (
	"yak/backend/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (apiVX *ApiV1) registerBoardsHandlers(router fiber.Router) {
	group := router.Group("/projects/:pid/boards")
	group.Get("/", apiVX.getBoards)
	group.Post("/", apiVX.createBoard)
	group.Get("/:bid", apiVX.getBoard)
	group.Put("/:bid", apiVX.updateBoard)
	group.Delete("/:bid", apiVX.deleteBoard)
}

func (apiVX *ApiV1) getBoards(ctx *fiber.Ctx) error {
	implementMe()
	boards := make([]models.Board, 0)
	return ctx.JSON(boards)
}

func (apiVX *ApiV1) getBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}

func (apiVX *ApiV1) createBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}

func (apiVX *ApiV1) updateBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}

func (apiVX *ApiV1) deleteBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}
