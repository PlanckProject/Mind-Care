package config

import (
	"os"
	"path"

	"github.com/PlanckProject/go-commons/config"
	"github.com/PlanckProject/go-commons/db/mongo"
	"github.com/PlanckProject/go-commons/logger"
)

type Configuration struct {
	App    AppConfig     `mapstructure:"app"`
	Mongo  mongo.Config  `mapstructure:"mongo"`
	Server HttpServer    `mapstructure:"server"`
	Logger logger.Config `mapstructure:"logger"`
	Maps   MapsConfig    `mapstructure:"maps"`
}

func New(configDir string) *Configuration {
	cfg := &Configuration{}
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	config.Parse(cfg, path.Join(configDir, env, "server.yml"))
	return cfg
}

type HttpServer struct {
	Port string `mapstructure:"port"`
}

type AppConfig struct {
	MaxQueryLimit      int64   `mapstructure:"max_query_limit"`
	MaxDistanceDefault float64 `mapstructure:"max_distance_default"`
}

type MapsConfig struct {
	LocationQueryPlaceholder  string       `mapstructure:"location_query_placeholder"`
	ProviderAPIKeyPlaceholder string       `mapstructure:"provider_api_key_placeholder"`
	OSM                       MapsProvider `mapstructure:"osm"`
	BingMaps                  MapsProvider `mapstructure:"bing_maps"`
}

type MapsProvider struct {
	Key     string `mapstructure:"key"`
	URL     string `mapstructure:"url"`
	Enabled bool   `mapstructure:"enabled"`
}
