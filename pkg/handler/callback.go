package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salihkemaloglu/todo/pkg/service"
)

func (h *Handler) Callback(ctx *gin.Context) {
	service.Callback(h.config)
	ctx.JSON(http.StatusOK, "response")
}
