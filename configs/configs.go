package configs

import (
	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
	"go.uber.org/dig"
)

type (
	configurations struct {
		Server server `toml:"server"`
	}

	server struct {
		Host string `toml:"host" validate:"required"`
		Port string `toml:"port" validate:"required"`
		Mode string `toml:"mode"`
	}
)

var Configurations configurations

const configFileLocation string = "./configs/configs.toml"

type LoadConfigsDeps struct {
	dig.In

	Validator *validator.Validate `name:"Validator"`
}

// LoadConfigs loads the configs file into the Go struct and validates it.
// Any changes to the configuration data or format('json', 'env') has to be made here.
func LoadConfigs(deps LoadConfigsDeps) error {
	if _, err := toml.DecodeFile(configFileLocation, &Configurations); err != nil {
		return err
	}

	if err := deps.Validator.Struct(&Configurations); err != nil {
		return err
	}

	return nil
}
