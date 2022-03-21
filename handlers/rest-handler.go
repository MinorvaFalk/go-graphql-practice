package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	r *gin.Engine
}

func NewRestHandler(r *gin.Engine) RestHandler {
	handler := RestHandler{r}

	return handler
}

func (rh RestHandler) SetupHandler() {
	rh.r.GET("/rest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "running",
		})
	})
}
