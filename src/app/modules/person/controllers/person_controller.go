package controllers

import (
	"log"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			dtos.BaseResponse{
				Title:   "Validation Error",
				Message: "'id' can only be numeric.",
			},
		)
		return
	}

	person, err := ctrl.personService.GetById(uint(id))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			dtos.BaseResponse{
				Title:   "Not Found",
				Message: "Person not found.",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		person,
	)
}

func (ctrl *PersonContoller) Create(c *gin.Context) {
	var requestBody personDtos.PersonCreateDto

	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			dtos.BaseResponse{
				Title:   "Validation Error",
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
				Title:   "Server Error",
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
