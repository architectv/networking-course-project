package v1

import (
	"fmt"
	"yak/backend/pkg/models"

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

func getUserId(ctx *fiber.Ctx) (string, error) {
	return "5fa8780a58e7c68e0a706032", nil  // ivan petrov
}

func (apiVX *ApiV1) getProjects(ctx *fiber.Ctx) error {
	// implementMe()
	// projects := make([]models.Project, 0)
	userId, err := getUserId(ctx)
	if err != nil {
		return err
	}

	projects, err := apiVX.services.Project.GetAll(userId)
	if err != nil {
		return err
	}
	return ctx.JSON(projects)
}

func (apiVX *ApiV1) getProject(ctx *fiber.Ctx) error {
	// implementMe()
	// project := models.Project{}
	userId, err := getUserId(ctx)
	if err != nil {
		return err
	}
	
	projectId := ctx.Params("pid")
	fmt.Println(projectId)
	project, err := apiVX.services.Project.GetById(userId, projectId)
	if err != nil {
		return err
	}
	return ctx.JSON(project)
}

func (apiVX *ApiV1) createProject(ctx *fiber.Ctx) error {
	// implementMe()
	// project := models.Project{}
	userId, err := getUserId(ctx)
	if err != nil {
		return err
	}

	var project models.Project
	if err := ctx.BodyParser(&project); err != nil {
		return err
	}
	id, err := apiVX.services.Project.Create(userId, project)
	if err != nil {
		return err
	}

	return ctx.JSON(map[string]interface{}{
		"id": id,
	})
}

func (apiVX *ApiV1) updateProject(ctx *fiber.Ctx) error {
	// implementMe()
	// project := models.Project{}
	userId, err := getUserId(ctx)
	if err != nil {
		return err
	}
	
	projectId := ctx.Params("pid")
	var project models.Project
	if err := ctx.BodyParser(&project); err != nil {
		return err
	}

	if err := apiVX.services.Project.Update(userId, projectId, project); err != nil {
		return err
	}
	return ctx.JSON(err)
}

func (apiVX *ApiV1) deleteProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}
