package v1

import (
	"yak/backend/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (apiVX *ApiV1) registerUsersHandlers(router fiber.Router) {
	group := router.Group("/users")
	group.Get("/", apiVX.getUsers)
	group.Post("/", apiVX.createUser)
	group.Get("/:uid", apiVX.getUser)
	group.Get("/login", apiVX.loginUser)
	group.Get("/logout", apiVX.logoutUser)
}

func (apiVX *ApiV1) getUsers(ctx *fiber.Ctx) error {
	// implementMe()
	// users := make([]models.User, 0)
	users, err := apiVX.services.User.GetAll()
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}

func (apiVX *ApiV1) getUser(ctx *fiber.Ctx) error {
	implementMe()
	user := models.User{}
	return ctx.JSON(user)
}

func (apiVX *ApiV1) createUser(ctx *fiber.Ctx) error {
	implementMe()
	user := models.User{}
	return ctx.JSON(user)
}

func (apiVX *ApiV1) loginUser(ctx *fiber.Ctx) error {
	implementMe()
	user := models.User{}
	return ctx.JSON(user)
}

func (apiVX *ApiV1) logoutUser(ctx *fiber.Ctx) error {
	implementMe()
	return ctx.Send([]byte{})
}
