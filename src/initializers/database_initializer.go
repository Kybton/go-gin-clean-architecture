package initializers

import (
	"fmt"
	"log"

	personModels "github.com/kybton/go-gin-clean-architecture/src/app/modules/person/models"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

type InitDbDep struct {
	dig.In

	DB *gorm.DB `name:"DatabaseConnection"`
}

func InitDatabase(deps InitDbDep) {
	fmt.Println("RUNNING DATABASE MIGRATE")

	log.Println(deps.DB)

	deps.DB.AutoMigrate(
		&personModels.Person{},
	)
}
