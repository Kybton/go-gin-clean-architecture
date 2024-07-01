package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kybton/go-gin-clean-architecture/src/app/controllers"
	"go.uber.org/dig"
)

type RootRouter struct {
	RootController controllers.RootController
}

// Dependencies struct for the root router
type RootRouterDeps struct {
	dig.In

	/// Dependencies for the root router
	RootController controllers.RootController `name:"RootController"`
}

func NewRootRouter(deps RootRouterDeps) RootRouter {
	return RootRouter{
		RootController: deps.RootController,
	}
}

// AddRoutes method create a new router group then add the routes and handlers to it.
func (r *RootRouter) AddRoutes(rg *gin.RouterGroup) {
	root := rg.Group("/")

	// Root server route
	root.GET("/", r.RootController.Root)
}
