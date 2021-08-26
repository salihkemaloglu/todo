package handler

import "github.com/salihkemaloglu/todo/pkg/config"

type Handler struct {
	config *config.Config
}

func NewHandler(conf *config.Config) *Handler {
	return &Handler{config: conf}
}
