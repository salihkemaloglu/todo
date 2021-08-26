package config

import (
	"github.com/salihkemaloglu/todo/pkg/agenc"
	"github.com/salihkemaloglu/todo/pkg/config"
)

func LoadConfig(filename string) *config.Config {
	var config config.Config
	err := agenc.ReadYAMLFile(filename, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
