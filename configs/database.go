package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(
		mysql.Open(
			Configurations.RDB.Username+":"+Configurations.RDB.Password+
				"@tcp("+Configurations.RDB.Host+":"+Configurations.RDB.Port+")/"+
				Configurations.RDB.Database+"?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{},
	)

	if err != nil {
		fmt.Println("gorm Db connection ", err)
		return nil, err
	}

	return db, err
}
