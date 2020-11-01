package v1

import "github.com/gofiber/fiber/v2"
import "yak/backend/pkg/handlers/api/v1/boards"
import "yak/backend/pkg/handlers/api/v1/lists"
import "yak/backend/pkg/handlers/api/v1/projects"
import "yak/backend/pkg/handlers/api/v1/tasks"
import "yak/backend/pkg/handlers/api/v1/users"

func RegisterHandlers(group fiber.Router) {
	v1_boards := group.Group("/boards")
	boards.RegisterHandlers(v1_boards)
	v1_lists := group.Group("/lists")
	lists.RegisterHandlers(v1_lists)
	v1_projects := group.Group("/projects")
	projects.RegisterHandlers(v1_projects)
	v1_tasks := group.Group("/tasks")
	tasks.RegisterHandlers(v1_tasks)
	v1_users := group.Group("/users")
	users.RegisterHandlers(v1_users)
}
