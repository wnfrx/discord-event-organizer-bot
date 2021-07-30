package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *restHandler) testWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Welcome!",
	})
}
