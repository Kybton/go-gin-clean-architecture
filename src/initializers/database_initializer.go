package initializers

import (
	"fmt"

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

	deps.DB.AutoMigrate(
		&personModels.Person{},
	)
}
