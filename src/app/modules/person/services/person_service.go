package services

import (
	"log"

	"github.com/kybton/go-gin-clean-architecture/src/app/modules/person/models"
	"github.com/kybton/go-gin-clean-architecture/src/app/modules/person/repositories"
	"go.uber.org/dig"
)

type PersonService struct {
	personRepository repositories.PersonRepsoitory
}

type PersonServiceDeps struct {
	dig.In

	PersonRepository repositories.PersonRepsoitory `name:"PersonRepository"`
}

func NewPersonService(deps PersonServiceDeps) PersonService {
	return PersonService{
		personRepository: deps.PersonRepository,
	}
}

func (ps *PersonService) List() {

}

func (ps *PersonService) GetById() {

}

func (ps *PersonService) Create(person *models.Person) (*models.Person, error) {
	_, err := ps.personRepository.Create(person)

	if err != nil {
		log.Fatal("Error occured")
		return person, err
	}

	return person, nil
}

func (ps *PersonService) Delete() {

}
