package infra

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/kybton/go-gin-clean-architecture/src/app/controllers"
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
			Construstor: root.NewRootRouter,
			Token:       "RootRouter",
		},
		{
			Construstor: controllers.NewRootController,
			Token:       "RootController",
		},
		{
			Construstor: validator.New,
			Token:       "Validator",
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
