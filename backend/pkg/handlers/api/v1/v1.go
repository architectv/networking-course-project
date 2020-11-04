package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"runtime"
)

func implementMe() {
	pc, fn, line, _ := runtime.Caller(1)
	fmt.Printf("Implement me in %s[%s:%d]\n", runtime.FuncForPC(pc).Name(), fn, line)
}

func RegisterHandlers(router fiber.Router) {
	v1 := router.Group("/v1")
	registerBoardsHandlers(v1)
	registerListsHandlers(v1)
	registerProjectsHandlers(v1)
	registerTasksHandlers(v1)
	registerUsersHandlers(v1)
}
