package main

import (
	"github.com/salihkemaloglu/todo/pkg/agenc"
	"github.com/salihkemaloglu/todo/pkg/command"
	"github.com/salihkemaloglu/todo/pkg/config"
	"github.com/spf13/cobra"
)

func main() {
	var configFile string // config file path
	var config config.Config

	loadConfig := func() {
		err := agenc.ReadYAMLFile(configFile, &config)
		if err != nil {
			panic(err)
		}
	}
	// initialize  config file options
	cobra.OnInitialize(loadConfig)
	var cmd = &cobra.Command{
		Use:   "todo",
		Short: "todo operations",
	}

	// config files with paths (--config="path/config.yml")
	cmd.PersistentFlags().StringVar(&configFile, "config", "config.yml", "config file")

	cmd.AddCommand(
		command.NewAPIRun(&config),
		command.NewConsumerRun(&config),
	)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
