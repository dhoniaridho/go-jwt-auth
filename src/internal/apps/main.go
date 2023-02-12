package internal

import (
	"api/src/internal/apps/auth"
	"api/src/internal/apps/todos"
	"api/src/internal/apps/users"

	"github.com/gin-gonic/gin"
)

type Router struct {
	URL   string
	Setup func(r *gin.RouterGroup) *gin.RouterGroup
}

func Setup(r *gin.Engine) *gin.Engine {

	routes := []Router{
		{
			URL:   "users",
			Setup: users.SetupRoutes,
		},
		{
			URL:   "todos",
			Setup: todos.SetupRoutes,
		},
		{
			URL:   "auth",
			Setup: auth.SetupRoutes,
		},
	}

	for _, route := range routes {

		currentRoute := r.Group(route.URL)
		route.Setup(currentRoute)
	}

	return r

}
