package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kybton/go-gin-clean-architecture/src/app/modules/person/controllers"
	"go.uber.org/dig"
)

type PersonRouter struct {
	personController controllers.PersonContoller
}

type PersonRouterDeps struct {
	dig.In

	PersonController controllers.PersonContoller `name:"PersonController"`
}

func NewPersonRouter(deps PersonRouterDeps) PersonRouter {
	return PersonRouter{
		personController: deps.PersonController,
	}
}

func (r *PersonRouter) AddRoutes(rg *gin.RouterGroup) {
	person := rg.Group("/person")

	person.GET("/", r.personController.Index)
	person.GET("/:id", r.personController.GetById)
	person.POST("/", r.personController.Create)
	person.DELETE("/:id", r.personController.Delete)
}
