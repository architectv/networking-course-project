package v1

import (
	"net/http"
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
	userId, err := apiVX.getUserId(ctx)
	if err != nil {
		return err
	}

	projectId := ctx.Params("pid")
	if projectId == "" {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	boards, err := apiVX.services.Board.GetAll(userId, projectId)
	if err != nil {
		return err
	}

	return ctx.JSON(boards)
}

func (apiVX *ApiV1) getBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}

func (apiVX *ApiV1) createBoard(ctx *fiber.Ctx) error {
	userId, err := apiVX.getUserId(ctx)
	if err != nil {
		return err
	}

	projectId := ctx.Params("pid")
	if projectId == "" {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	var input models.Board
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	boardId, err := apiVX.services.Board.Create(userId, projectId, input)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"bid": boardId,
	})
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
