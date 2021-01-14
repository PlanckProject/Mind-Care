package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerHealthCheck(g *gin.Engine) {
	g.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
