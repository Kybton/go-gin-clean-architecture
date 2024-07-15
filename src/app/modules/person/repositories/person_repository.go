package repositories

import (
	"errors"
	"log"
	"time"

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

func (pr *PersonRepsoitory) Get(id uint) (*personModels.Person, error) {
	var person *personModels.Person

	result := pr.db.Find(&person, id)

	if result.Error != nil {
		return person, result.Error
	}

	if result.RowsAffected == 0 {
		return person, errors.New("not found")
	}

	return person, nil
}

func (pr *PersonRepsoitory) Create(person *personModels.Person) (*personModels.Person, error) {
	now := time.Now().UTC()
	person.CreatedAt = &now
	person.UpdatedAt = &now
	person.DeletedAt = nil

	result := pr.db.Create(person)

	if result.Error != nil {
		log.Fatal(result.Error)
		return person, result.Error
	}

	return person, nil
}
