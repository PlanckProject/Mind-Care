package config

import (
	"github.com/PlanckProject/go-commons/config"
	"github.com/PlanckProject/go-commons/db/mongo"
	"github.com/PlanckProject/go-commons/logger"
)

type Configuration struct {
	Mongo  mongo.Config  `mapstructure:"mongo"`
	Server HttpServer    `mapstructure:"server"`
	Logger logger.Config `mapstructure:"logger"`
}

func New() *Configuration {
	cfg := &Configuration{}
	config.Parse(cfg, "")
	return cfg
}

type HttpServer struct {
	Port string
}
