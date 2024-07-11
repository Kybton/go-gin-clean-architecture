package main

import (
	"log"
	"os"

	"github.com/kybton/go-gin-clean-architecture/configs"
	"github.com/kybton/go-gin-clean-architecture/src/app"
	"github.com/kybton/go-gin-clean-architecture/src/infra"
	"github.com/kybton/go-gin-clean-architecture/src/initializers"
	"go.uber.org/dig"
)

var diContainer *dig.Container

func init() {
	// Loading configs on the application starts up
	var err error

	diContainer, err = infra.DependencyInjector()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err = diContainer.Invoke(configs.LoadConfigs); err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	if err = diContainer.Invoke(initializers.InitDatabase); err != nil {
		log.Fatal(err)
		os.Exit(3)
	}
}

func main() {
	if err := diContainer.Invoke(app.NewServer); err != nil {
		log.Fatal(err)
		os.Exit(3)
	}
}
