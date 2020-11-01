package v1

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/models"

func registerBoardsHandlers(router fiber.Router) {
	group := router.Group("/projects/:pid/boards")
	group.Get("/", getBoards)
	group.Post("/", createBoard)
	group.Get("/:bid", getBoard)
	group.Put("/:bid", updateBoard)
	group.Delete("/:bid", deleteBoard)
}

func getBoards(ctx *fiber.Ctx) error {
	implementMe()
	boards := make([]models.Board, 0)
	return ctx.JSON(boards)
}

func getBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}

func createBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}

func updateBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}

func deleteBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}
