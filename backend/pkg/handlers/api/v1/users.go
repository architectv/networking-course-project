package v1

import (
	"errors"
	"strings"
	"yak/backend/pkg/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func (apiVX *ApiV1) registerUsersHandlers(router fiber.Router) {
	group := router.Group("/users")
	// group.Get("/", apiVX.getUsers)
	// group.Post("/", apiVX.createUser)
	// group.Get("/:uid", apiVX.getUser)
	group.Post("/signup", apiVX.signUp)
	group.Post("/signin", apiVX.signIn)
	group.Get("/signout", apiVX.userIdentity, apiVX.signOut)
}

// func (apiVX *ApiV1) getUsers(ctx *fiber.Ctx) error {
// 	users, err := apiVX.services.User.GetAll()
// 	if err != nil {
// 		return err
// 	}
// 	return ctx.JSON(users)
// }

// type signInInput struct {
// 	Nickname string `json:"username" valid:"length(3|32)"`
// 	Password string `json:"password" valid:"length(6|32)"`
// }

// sigIn function
func (apiVX *ApiV1) signIn(ctx *fiber.Ctx) error {
	response := &models.ApiResponse{}
	type signInInput struct {
		Nickname string `json:"username" valid:"length(3|32)"`
		Password string `json:"password" valid:"length(6|32)"`
	}
	var input signInInput

	if err := ctx.BodyParser(&input); err != nil {
		response.Error(fiber.StatusBadRequest, err.Error())
		return Send(ctx, response)
	}
	if _, err := govalidator.ValidateStruct(input); err != nil {
		response.Error(fiber.StatusBadRequest, err.Error())
		return Send(ctx, response)
	}

	response = apiVX.services.User.GenerateToken(ctx.Context(), input.Nickname, input.Password)
	return Send(ctx, response)
}

// sigUp function
func (apiVX *ApiV1) signUp(ctx *fiber.Ctx) error {
	response := &models.ApiResponse{}
	input := &models.User{}

	if err := ctx.BodyParser(input); err != nil {
		response.Error(fiber.StatusBadRequest, err.Error())
		return Send(ctx, response)
	}
	if _, err := govalidator.ValidateStruct(input); err != nil {
		response.Error(fiber.StatusBadRequest, err.Error())
		return Send(ctx, response)
	}
	response = apiVX.services.User.Create(ctx.Context(), input)
	return Send(ctx, response)
}

func (apiVX *ApiV1) signOut(ctx *fiber.Ctx) error {
	response := &models.ApiResponse{}
	_, err := apiVX.getUserId(ctx)
	if err != nil {
		response.Error(fiber.StatusUnauthorized, err.Error())
		return Send(ctx, response)
	}

	token, err := apiVX.getToken(ctx)
	if err != nil {
		response.Error(fiber.StatusUnauthorized, err.Error())
		return Send(ctx, response)
	}

	response = apiVX.services.User.SignOut(ctx.Context(), token)
	return Send(ctx, response)
}

const (
	authorizationHeader = "Authorization"
	userCtx             = "_id"
)

func (apiVX *ApiV1) getToken(ctx *fiber.Ctx) (string, error) {
	header := ctx.Get(authorizationHeader)
	if header == "" {
		return "", errors.New("Empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("Invalid token")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("Empty token")
	}

	return headerParts[1], nil
}

func (apiVX *ApiV1) userIdentity(ctx *fiber.Ctx) error {
	response := &models.ApiResponse{}

	token, err := apiVX.getToken(ctx)
	if err != nil {
		response.Error(fiber.StatusUnauthorized, err.Error())
		return Send(ctx, response)
	}

	userId, err := apiVX.services.User.ParseToken(ctx.Context(), token)
	if err != nil {
		response.Error(fiber.StatusUnauthorized, err.Error())
		return Send(ctx, response)
	}

	ctx.Request().Header.Set(userCtx, userId)
	return ctx.Next()
	// response.Set(fiber.StatusOK, "OK", fiber.Map{userCtx: userId})
	// return Send(ctx, response)
}

func (apiVX *ApiV1) getUserId(ctx *fiber.Ctx) (string, error) {
	id := ctx.Get(userCtx)
	if id == "" {
		return "", errors.New("user id not found")
	}

	return id, nil
}
