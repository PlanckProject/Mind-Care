package config

import (
	"github.com/PlankProject/go-commons/config"
	"github.com/PlankProject/go-commons/db/mongo"
	"github.com/PlankProject/go-commons/logger"
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
