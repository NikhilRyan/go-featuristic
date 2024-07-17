package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	ServerPort string
	DBConfig   map[string]string
	CacheHost  string
	CachePort  string
	BaseURL    string
}

var DBConfig = map[string]string{
	"driver":      "postgres",
	"conn_string": "host=localhost user=postgres password=password dbname=gofeaturistic port=5432 sslmode=disable",
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	config.BaseURL = "http://localhost:" + config.ServerPort
	return &config, nil
}

func GetMainDSN() string {
	return DBConfig["conn_string"]
}

func IsProd() bool {
	return os.Getenv("ENV") == "prod"
}

func GetRedisAddr() string {
	return os.Getenv("REDIS_ADDR")
}
