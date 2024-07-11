package infra

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/kybton/go-gin-clean-architecture/configs"
	"github.com/kybton/go-gin-clean-architecture/src/app/controllers"
	personController "github.com/kybton/go-gin-clean-architecture/src/app/modules/person/controllers"
	personRepository "github.com/kybton/go-gin-clean-architecture/src/app/modules/person/repositories"
	personRouterV1 "github.com/kybton/go-gin-clean-architecture/src/app/modules/person/routers/v1"
	personService "github.com/kybton/go-gin-clean-architecture/src/app/modules/person/services"
	root "github.com/kybton/go-gin-clean-architecture/src/app/routers"
	"go.uber.org/dig"
)

type dependency struct {
	Construstor interface{}
	Token       string
}

func DependencyInjector() (*dig.Container, error) {
	deps := []dependency{
		{
			Construstor: configs.DbConnection,
			Token:       "DatabaseConnection",
		},
		{
			Construstor: validator.New,
			Token:       "Validator",
		},
		{
			Construstor: root.NewRootRouter,
			Token:       "RootRouter",
		},
		{
			Construstor: controllers.NewRootController,
			Token:       "RootController",
		},
		{
			Construstor: personRouterV1.NewPersonRouter,
			Token:       "PersonRouter",
		},
		{
			Construstor: personController.NewPersonController,
			Token:       "PersonController",
		},
		{
			Construstor: personService.NewPersonService,
			Token:       "PersonService",
		},
		{
			Construstor: personRepository.NewPersonRepository,
			Token:       "PersonRepository",
		},
	}

	container := dig.New()

	for _, dep := range deps {
		if err := container.Provide(
			dep.Construstor,
			dig.Name(dep.Token),
		); err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return container, nil
}
