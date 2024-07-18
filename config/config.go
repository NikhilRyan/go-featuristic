package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Config contains database configuration properties
type Config struct {
	Driver     string
	Host       string
	Port       string
	User       string
	Password   string
	DBName     string
	SSLMode    string // For Postgres
	Charset    string // For MySQL
	BaseURL    string
	ServerPort string
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

// GetMainDSN constructs the DSN based on the loaded configuration
func GetMainDSN(cfg Config) string {
	switch cfg.Driver {
	case "postgres":
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset)
	default:
		return ""
	}
}

func IsProd() bool {
	return os.Getenv("ENV") == "prod"
}

func GetRedisAddr() string {
	return os.Getenv("REDIS_ADDR")
}
