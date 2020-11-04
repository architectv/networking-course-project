package v1

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/models"

func registerUsersHandlers(router fiber.Router) {
	group := router.Group("/users")
	group.Get("/", getUsers)
	group.Post("/", createUser)
	group.Get("/:uid", getUser)
	group.Get("/login", loginUser)
	group.Get("/logout", logoutUser)
}

func getUsers(ctx *fiber.Ctx) error {
	implementMe()
	users := make([]models.User, 0)
	return ctx.JSON(users)
}

func getUser(ctx *fiber.Ctx) error {
	implementMe()
	user := models.User{}
	return ctx.JSON(user)
}

func createUser(ctx *fiber.Ctx) error {
	implementMe()
	user := models.User{}
	return ctx.JSON(user)
}

func loginUser(ctx *fiber.Ctx) error {
	implementMe()
	user := models.User{}
	return ctx.JSON(user)
}

func logoutUser(ctx *fiber.Ctx) error {
	implementMe()
	return ctx.Send([]byte{})
}
