package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salihkemaloglu/todo/pkg/model"
	"github.com/salihkemaloglu/todo/pkg/service"
)

func (h *Handler) Callback(ctx *gin.Context) {
	var o model.Object
	if err := ctx.ShouldBindJSON(&o); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := service.Send(o, h.config)
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}
