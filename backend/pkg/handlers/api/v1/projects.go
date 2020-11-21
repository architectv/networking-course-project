package v1

import (
	"errors"
	"strconv"
	"yak/backend/pkg/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func (apiVX *ApiV1) registerProjectsHandlers(router fiber.Router) {
	group := router.Group("/projects")
	group.Get("/", apiVX.getProjects)
	group.Post("/", apiVX.createProject)
	group.Get("/:pid", apiVX.getProject)
	group.Put("/:pid", apiVX.updateProject)
	group.Delete("/:pid", apiVX.deleteProject)
}

func (apiVX *ApiV1) createProject(ctx *fiber.Ctx) error {
	// implementMe()
	// project := models.Project{}
	response := &models.ApiResponse{}
	userId, err := apiVX.getUserId(ctx)
	if err != nil {
		return err
	}

	project := &models.Project{}

	if err := ctx.BodyParser(project); err != nil {
		response.Error(fiber.StatusBadRequest, err.Error())
		return err
	}

	if _, err := govalidator.ValidateStruct(project); err != nil {
		response.Error(fiber.StatusBadRequest, err.Error())
		return Send(ctx, response)
	}

	response = apiVX.services.Project.Create(userId, project)
	return Send(ctx, response)
}

func (apiVX *ApiV1) getProjects(ctx *fiber.Ctx) error {
	// implementMe()
	// projects := make([]models.Project, 0)
	userId, err := apiVX.getUserId(ctx)
	if err != nil {
		return err
	}

	response := apiVX.services.Project.GetAll(userId)
	return Send(ctx, response)
}

func (apiVX *ApiV1) getProject(ctx *fiber.Ctx) error {
	// implementMe()
	// project := models.Project{}
	userId, err := apiVX.getUserId(ctx)
	if err != nil {
		return err
	}

	projectId := ctx.Params("pid")
	intProjectId, err := strconv.ParseInt(projectId, 10, 64) // TODO сделать нормальный перевод строки в число
	if err != nil {
		return errors.New("project id is of invalid type")
	}

	response := apiVX.services.Project.GetById(userId, int(intProjectId))
	return Send(ctx, response)
}

func (apiVX *ApiV1) updateProject(ctx *fiber.Ctx) error {
	implementMe()
	// project := models.Project{}
	// userId, err := apiVX.getUserId(ctx)
	// if err != nil {
	// 	return err
	// }

	// projectId := ctx.Params("pid")
	// var project models.Project
	// if err := ctx.BodyParser(&project); err != nil {
	// 	return err
	// }

	// if err := apiVX.services.Project.Update(userId, projectId, project); err != nil {
	// 	return err
	// }
	// return ctx.JSON(err)
	return nil
}

func (apiVX *ApiV1) deleteProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}
