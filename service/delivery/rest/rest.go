package rest

import (
	"github.com/gin-gonic/gin"
)

type restHandler struct {
	router *gin.Engine
}

func NewRestHandler(
	r *gin.Engine,
) *restHandler {
	return &restHandler{
		router: r,
	}
}

func (h *restHandler) RegisterRoutes() {
	h.router.GET("/", h.testWelcome)
}
