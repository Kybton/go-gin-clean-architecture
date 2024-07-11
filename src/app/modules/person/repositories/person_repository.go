package repositories

import (
	"log"

	personModels "github.com/kybton/go-gin-clean-architecture/src/app/modules/person/models"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type PersonRepsoitory struct {
	db *gorm.DB
}

type PersonRepositoryDeps struct {
	dig.In

	DB *gorm.DB `name:"DatabaseConnection"`
}

func NewPersonRepository(deps PersonRepositoryDeps) PersonRepsoitory {
	return PersonRepsoitory{
		db: deps.DB,
	}
}

func (pr *PersonRepsoitory) Create(person *personModels.Person) (*personModels.Person, error) {
	result := pr.db.Create(person)

	if result.Error != nil {
		log.Fatal(result.Error)
		return person, result.Error
	}

	return person, nil
}
