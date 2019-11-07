package config

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

const (
	ConfigFilePath = "./config/config.toml"
)

var (
	Conf = new()
)

/**
 * Set parse Toml file
 */
func new() *toml.Tree {
	config, err := toml.LoadFile(ConfigFilePath)

	if err != nil {
		fmt.Println("TomlErr ", err.Error())
	}

	return config
}
