package server

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/PlankProject/Mental-Wellbeing-Resources/api/config"
	"github.com/PlankProject/Mental-Wellbeing-Resources/api/pkg/io/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(New), http.Module)

func New(cfg *config.Configuration, logWriter io.Writer) *gin.Engine {
	g := gin.New()
	g.Use(gin.LoggerWithConfig(getGinLoggerConfig(cfg, logWriter)), gin.RecoveryWithWriter(logWriter))
	return g
}

func getGinLoggerConfig(cfg *config.Configuration, writer io.Writer) gin.LoggerConfig {
	loggerConfig := gin.LoggerConfig{}

	// Formatter
	switch strings.ToLower(cfg.Logger.Format) {
	case "json":
		fallthrough
	default:
		loggerConfig.Formatter = formatter
	}

	// Output
	loggerConfig.Output = writer

	return loggerConfig
}

func formatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf(`{"src":"GIN","client_ip":"%s","timestamp":"%s","request.method":"%s","request.path":"%s","request.proto":"%s","request.status.code":%d,"latency":"%s","user_agent":"%s","error":"%s"}`,
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
