package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestHandler struct{}

func NewRestHandler(r *gin.Engine) {
	// handler := RestHandler{}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "running",
		})
	})
}
