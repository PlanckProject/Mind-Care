package main

import (
	"net/http"

	"github.com/PlankProject/Mental-Wellbeing-Resources/api/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(fx.Provide(config.New))
	app.Run()
	g := gin.New()
	g.GET("*", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{})
	})
}
