package internal

import (
	"github.com/gin-gonic/gin"
	"server/tools/wrapper"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	h.initRoutes(router)
	return router
}

func (h *Handler) initRoutes(router *gin.Engine) {
	router.Use(wrapper.RouteWrapper())
	router.GET("/values/all", h.GetValues)
	router.GET("/values/current", h.GetCurrent)
	router.POST("/values", h.Create)
}
