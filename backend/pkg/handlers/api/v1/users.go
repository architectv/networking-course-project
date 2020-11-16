package v1

import (
	"errors"
	"net/http"
	"strings"
	"yak/backend/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
	users, err := apiVX.services.User.GetAll()
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}

type signInInput struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

func (apiVX *ApiV1) getUser(ctx *fiber.Ctx) error {
	var input signInInput

	if err := ctx.BodyParser(&input); err != nil {
		logrus.Println("Bad Request")
		return ctx.SendStatus(http.StatusBadRequest)
	}
	if input.Username == "" || input.Password == "" {
		logrus.Println("Bad Request")
		return ctx.SendStatus(http.StatusBadRequest)
	}

	token, err := apiVX.services.User.GenerateToken(input.Username, input.Password)
	if err != nil {
		logrus.Println("Internal Server Error")
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func (apiVX *ApiV1) createUser(ctx *fiber.Ctx) error {
	var input models.User

	if err := ctx.BodyParser(&input); err != nil {
		logrus.Println("Bad Request")
		return ctx.SendStatus(http.StatusBadRequest)
	}

	id, err := apiVX.services.User.Create(input)
	if err != nil {
		logrus.Println("Internal Server Error")
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"_id": id,
	})
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

const (
	authorizationHeader = "Authorization"
	userCtx             = "_id"
)

func (apiVX *ApiV1) userIdentity(ctx *fiber.Ctx) error {
	header := ctx.Get(authorizationHeader)
	if header == "" {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	if len(headerParts[1]) == 0 {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	userId, err := apiVX.services.User.ParseToken(headerParts[1])
	if err != nil {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	ctx.Request().Header.Set(userCtx, userId)
	return ctx.Next()
}

func (apiVX *ApiV1) getUserId(ctx *fiber.Ctx) (string, error) {
	id := ctx.Get(userCtx)
	if id == "" {
		return "", errors.New("user id not found")
	}

	return id, nil
}
