package v1

import (
	"strconv"
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
	response := &models.ApiResponse{}
	userId, err := getUserId(ctx)
	if err != nil {
		response.Error(fiber.StatusInternalServerError, err.Error())
		return Send(ctx, response)
	}

	projectId, err := strconv.Atoi(ctx.Params("pid"))
	if err != nil || projectId == 0 {
		response.Error(fiber.StatusBadRequest, "Empty projectId")
		return Send(ctx, response)
	}

	boardId, err := strconv.Atoi(ctx.Params("bid"))
	if err != nil || boardId == 0 {
		response.Error(fiber.StatusBadRequest, "Invalid boardId")
		return Send(ctx, response)
	}

	listId, err := strconv.Atoi(ctx.Params("lid"))
	if err != nil || listId == 0 {
		response.Error(fiber.StatusBadRequest, "Invalid listId")
		return Send(ctx, response)
	}

	response = apiVX.services.Task.GetAll(userId, projectId, boardId, listId)
	return Send(ctx, response)
}

func (apiVX *ApiV1) getTask(ctx *fiber.Ctx) error {
	response := &models.ApiResponse{}
	userId, err := getUserId(ctx)
	if err != nil {
		response.Error(fiber.StatusInternalServerError, err.Error())
		return Send(ctx, response)
	}

	projectId, err := strconv.Atoi(ctx.Params("pid"))
	if err != nil || projectId == 0 {
		response.Error(fiber.StatusBadRequest, "Invalid projectId")
		return Send(ctx, response)
	}

	boardId, err := strconv.Atoi(ctx.Params("bid"))
	if err != nil || boardId == 0 {
		response.Error(fiber.StatusBadRequest, "Invalid boardId")
		return Send(ctx, response)
	}

	listId, err := strconv.Atoi(ctx.Params("lid"))
	if err != nil || listId == 0 {
		response.Error(fiber.StatusBadRequest, "Invalid listId")
		return Send(ctx, response)
	}

	taskId, err := strconv.Atoi(ctx.Params("tid"))
	if err != nil || taskId == 0 {
		response.Error(fiber.StatusBadRequest, "Invalid taskId")
		return Send(ctx, response)
	}

	response = apiVX.services.Task.GetById(userId, projectId, boardId,
		listId, taskId)
	return Send(ctx, response)
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
