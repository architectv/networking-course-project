package v1

import (
	"fmt"
	"runtime"
	"yak/backend/pkg/models"
	"yak/backend/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func implementMe() {
	pc, fn, line, _ := runtime.Caller(1)
	fmt.Printf("Implement me in %s[%s:%d]\n", runtime.FuncForPC(pc).Name(), fn, line)
}

type ApiV1 struct {
	services *services.Service
}

func NewApiV1(services *services.Service) *ApiV1 {
	return &ApiV1{services: services}
}

func (apiVX *ApiV1) RegisterHandlers(router fiber.Router) {
	v1 := router.Group("/v1", apiVX.userIdentity)
	apiVX.registerBoardsHandlers(v1)
	apiVX.registerListsHandlers(v1)
	apiVX.registerProjectPermsHandlers(v1)
	apiVX.registerProjectsHandlers(v1)
	apiVX.registerTasksHandlers(v1)

	auth := router.Group("/auth")
	apiVX.registerUsersHandlers(auth)
}

func Send(ctx *fiber.Ctx, r *models.ApiResponse) error {
	ctx.Status(r.Code)
	return ctx.JSON(r)
}
