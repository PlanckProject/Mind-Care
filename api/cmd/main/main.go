package main

import (
	"flag"

	"github.com/PlanckProject/Mind-Care/api/config"
	"github.com/PlanckProject/Mind-Care/api/pkg/core/logger"
	"github.com/PlanckProject/Mind-Care/api/pkg/io/http/server"
	"github.com/PlanckProject/Mind-Care/api/pkg/repo"
	"github.com/PlanckProject/Mind-Care/api/pkg/service"
	"go.uber.org/fx"
)

var configDir string

func init() {
	flag.StringVar(&configDir, "configDir", "config", "Relative path of the directory containing the config file")
	flag.Parse()
}

func main() {
	app := fx.New(
		fx.Provide(func() string { return configDir }),
		fx.Provide(config.New),
		logger.Module,
		repo.Module,
		service.Module,
		server.Module,
	)
	app.Run()
}
