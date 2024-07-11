package services

import (
	"log"

	"github.com/kybton/go-gin-clean-architecture/src/app/modules/person/models"
	"github.com/kybton/go-gin-clean-architecture/src/app/modules/person/repositories"
	"go.uber.org/dig"
)

type PersonService struct {
	PersonRepository repositories.PersonRepsoitory
}

type PersonServiceDeps struct {
	dig.In

	PersonRepository repositories.PersonRepsoitory `name:"PersonRepository"`
}

func NewPersonService(deps PersonServiceDeps) PersonService {
	return PersonService{
		PersonRepository: deps.PersonRepository,
	}
}

func (pr *PersonService) Create(person *models.Person) *models.Person {
	_, err := pr.PersonRepository.Create(person)

	if err != nil {
		log.Fatal("Error occured")
	}

	return person
}
