package handlers

import (
	"net/http"

	"github.com/PlanckProject/Mind-Care/api/pkg/core/utils"
	"github.com/PlanckProject/Mind-Care/api/pkg/repo"
	"github.com/gin-gonic/gin"
)

func registerHealthCheck(g *gin.Engine, repo repo.IServiceProvidersRepo) {
	g.GET("/health", func(c *gin.Context) {
		err := repo.Ping(c.Request.Context())
		db := "connected"
		responseCode := http.StatusOK
		if err != nil {
			db = "disconnected"
			responseCode = http.StatusServiceUnavailable
		}

		response := map[string]interface{}{
			"service":  "healthy",
			"database": db,
		}

		switch responseCode {
		case http.StatusServiceUnavailable:
			utils.Respond(c, responseCode, response, nil)
		default:
			utils.Respond(c, responseCode, response, nil)
		}
	})
}
