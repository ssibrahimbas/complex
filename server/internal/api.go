package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/entity"
	"server/tools/result"
)

func (h *Handler) GetValues(context *gin.Context) {
	values, err := h.service.GetAll()
	if err != nil {
		context.Error(err)
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, result.SuccessData("Values fetch success.", values))
	context.Abort()
}

func (h *Handler) GetCurrent(context *gin.Context) {
	values, err := h.service.GetCurrent()
	if err != nil {
		context.Error(err)
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, result.SuccessData("Values fetch success.", values))
	context.Abort()
}

func (h *Handler) Create(context *gin.Context) {
	var input *entity.Value
	context.BindJSON(&input)
	product, err := h.service.CreateValue(input)
	if err != nil {
		context.Error(err)
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, result.SuccessData("Product created success.", product))
	context.Abort()
}
