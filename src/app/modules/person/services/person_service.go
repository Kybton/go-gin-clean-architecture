package services

import (
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

func (ps *PersonService) GetById(id uint) (*models.Person, error) {
	return ps.personRepository.Get(id)
}

func (ps *PersonService) Create(person *models.Person) (*models.Person, error) {
	return ps.personRepository.Create(person)
}

func (ps *PersonService) Delete() {

}
