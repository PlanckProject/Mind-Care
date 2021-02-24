package utils

import (
	"strconv"

	"github.com/PlanckProject/go-commons/logger"
	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, status int, payload interface{}, err error) {
	if err != nil {
		logger.WithFields(map[string]interface{}{
			"http.request.method":  c.Request.Method,
			"http.request.uri":     c.Request.URL.RequestURI(),
			"http.response.status": strconv.Itoa(status),
			"client.ip":            c.ClientIP(),
		}).Warn(err.Error())

		c.JSON(status, map[string]interface{}{
			"error": err.Error(),
		})
	} else {
		if payload != nil {
			logger.WithFields(map[string]interface{}{
				"http.request.method":   c.Request.Method,
				"http.request.uri":      c.Request.URL.RequestURI(),
				"http.response.status":  strconv.Itoa(status),
				"client.ip":             c.ClientIP(),
				"http.response.payload": payload,
			}).Info("OK")

			c.JSON(status, map[string]interface{}{
				"message": "success",
				"data":    payload,
			})
		} else {
			logger.WithFields(map[string]interface{}{
				"http.request.method":  c.Request.Method,
				"http.request.uri":     c.Request.URL.RequestURI(),
				"http.response.status": strconv.Itoa(status),
				"client.ip":            c.ClientIP(),
			}).Info("OK")

			c.JSON(status, map[string]interface{}{"message": "success"})
		}
	}
	c.Abort()
}
