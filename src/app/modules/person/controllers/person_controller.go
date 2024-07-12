package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kybton/go-gin-clean-architecture/src/app/dtos"
	personDtos "github.com/kybton/go-gin-clean-architecture/src/app/modules/person/dtos"
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
	// TODO
}

func (ctrl *PersonContoller) GetById(c *gin.Context) {
	// TODO
}

func (ctrl *PersonContoller) Create(c *gin.Context) {
	var requestBody personDtos.PersonCreateDto

	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			dtos.BaseResponse{
				Message: "Invalid request body.",
			},
		)
		return
	}

	person, err := ctrl.personService.Create(&models.Person{
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
	})

	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			dtos.BaseResponse{
				Message: "Error occurs while creating person.",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		person,
	)
}

func (ctrl *PersonContoller) Delete(c *gin.Context) {
	// TODO
}
