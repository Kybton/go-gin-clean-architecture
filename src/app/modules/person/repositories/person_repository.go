package repositories

import (
	"log"

	personModels "github.com/kybton/go-gin-clean-architecture/src/app/modules/person/models"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type PersonRepsoitory struct {
	DB *gorm.DB
}

type PersonRepositoryDeps struct {
	dig.In

	DB *gorm.DB `name:"DatabaseConnection"`
}

func NewPersonRepository(deps PersonRepositoryDeps) PersonRepsoitory {
	return PersonRepsoitory{
		DB: deps.DB,
	}
}

func (r *PersonRepsoitory) Create(person *personModels.Person) (*personModels.Person, error) {
	result := r.DB.Create(person)

	if result.Error != nil {
		log.Fatal(result.Error)
		return person, result.Error
	}

	return person, nil
}
