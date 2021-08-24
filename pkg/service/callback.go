package service

import (
	"github.com/salihkemaloglu/todo/pkg/pgu"
	"github.com/salihkemaloglu/todo/pkg/repository"
	"github.com/salihkemaloglu/todo/pkg/util/config"
)

func Callback(config *config.Config) {
	node, err := pgu.MustOpen(config.Postgres.DatabaseURL)
	if err != nil {

	}

	repository.InsertTodo(1, node)

}
