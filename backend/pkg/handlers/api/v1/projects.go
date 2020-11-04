package v1

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/models"

func registerProjectsHandlers(router fiber.Router) {
	group := router.Group("/projects")
	group.Get("/", getProjects)
	group.Post("/", createProject)
	group.Get("/:pid", getProject)
	group.Put("/:pid", updateProject)
	group.Delete("/:pid", deleteProject)
}

func getProjects(ctx *fiber.Ctx) error {
	implementMe()
	projects := make([]models.Project, 0)
	return ctx.JSON(projects)
}

func getProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}

func createProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}

func updateProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}

func deleteProject(ctx *fiber.Ctx) error {
	implementMe()
	project := models.Project{}
	return ctx.JSON(project)
}
