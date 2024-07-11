package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kybton/go-gin-clean-architecture/src/app/modules/person/models"
	"github.com/kybton/go-gin-clean-architecture/src/app/modules/person/services"
	"go.uber.org/dig"
)

type PersonContoller struct {
	personService services.PersonService
}

type PersonContollerDeps struct {
	dig.In

	PersonService services.PersonService `name:"PersonService"`
}

func NewPersonController(deps PersonContollerDeps) PersonContoller {
	return PersonContoller{
		personService: deps.PersonService,
	}
}

func (ctrl *PersonContoller) Index(c *gin.Context) {

}

func (ctrl *PersonContoller) GetById(c *gin.Context) {

}

func (ctrl *PersonContoller) Create(c *gin.Context) {
	person := ctrl.personService.Create(&models.Person{
		FirstName: "First Name",
		LastName:  "LastName",
	})

	c.JSON(
		http.StatusOK,
		person,
	)
}
