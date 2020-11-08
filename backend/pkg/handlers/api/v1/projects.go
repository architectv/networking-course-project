package v1

import (
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

func (apiVX *ApiV1) getProjects(ctx *fiber.Ctx) error {
	implementMe()
	projects := make([]models.Project, 0)
	return ctx.JSON(projects)
}

func (apiVX *ApiV1) getProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}

func (apiVX *ApiV1) createProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}

func (apiVX *ApiV1) updateProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}

func (apiVX *ApiV1) deleteProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}
