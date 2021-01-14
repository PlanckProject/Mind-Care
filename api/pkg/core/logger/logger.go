package logger

import (
	"io"
	"os"

	"github.com/PlankProject/Mental-Wellbeing-Resources/api/config"
	"github.com/PlankProject/go-commons/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(rotatedWriter), fx.Invoke(initialize))

func rotatedWriter(cfg *config.Configuration) io.Writer {
	if cfg.Logger.Enabled {
		return logger.GetRotatedWriter(&cfg.Logger)
	}
	return os.Stdout
}

func initialize(cfg *config.Configuration, writer io.Writer) {
	logger.Configure(&cfg.Logger, writer)
}
