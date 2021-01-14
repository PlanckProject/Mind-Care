package main

import (
	"github.com/PlankProject/Mental-Wellbeing-Resources/api/config"
	"github.com/PlankProject/Mental-Wellbeing-Resources/api/pkg/core/logger"
	"github.com/PlankProject/Mental-Wellbeing-Resources/api/pkg/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(config.New),
		logger.Module,
		server.Module,
	)
	app.Run()
}
