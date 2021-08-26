package handler

import "github.com/salihkemaloglu/todo/pkg/util/config"

// Handler is endpoints
type Handler struct {
	config *config.Config
}

// NewHandler example
func NewHandler(conf *config.Config) *Handler {
	return &Handler{config: conf}
}
