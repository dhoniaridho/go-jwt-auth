package main

import (
	"api/internal/apps/todos"
	"api/internal/apps/users"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	type Router struct {
		URL   string
		Setup func(r *gin.RouterGroup) *gin.RouterGroup
	}

	routes := []Router{
		{
			URL:   "users",
			Setup: users.SetupRoutes,
		},
		{
			URL:   "todos",
			Setup: todos.SetupRoutes,
		},
	}

	for _, route := range routes {

		currentRoute := r.Group(route.URL)
		route.Setup(currentRoute)
	}

	r.Run()

}
