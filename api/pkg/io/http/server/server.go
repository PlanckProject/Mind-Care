package server

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/PlanckProject/Mind-Care/api/config"
	"github.com/PlanckProject/Mind-Care/api/pkg/io/http/handlers"
	"github.com/PlanckProject/go-commons/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(New), handlers.Module, fx.Invoke(start))

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

func start(g *gin.Engine, cfg *config.Configuration, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			if cfg.Server.Port == "" {
				cfg.Server.Port = "8080"
			}
			port := fmt.Sprintf(":%s", cfg.Server.Port)
			logger.Info("Server running on ", port)
			go func() {
				err := g.Run(port)
				if err != nil {
					logger.WithField("error", err).Error("Failed to start server")
					os.Exit(1)
				}
			}()
			return nil
		},
	})
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
