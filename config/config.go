package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	CacheHost  string `mapstructure:"CACHE_HOST"`
	CachePort  string `mapstructure:"CACHE_PORT"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

var DBConfig = map[string]string{
	"driver":      "postgres",
	"conn_string": "host=localhost user=postgres password=password dbname=gofeaturistic port=5432 sslmode=disable",
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func GetDSN(cfg Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
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
