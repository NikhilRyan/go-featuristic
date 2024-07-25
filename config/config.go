package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Config contains database configuration properties
type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	RedisAddr  string `mapstructure:"REDIS_ADDR"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	Driver     string `mapstructure:"DRIVER"`
	SSLMode    string `mapstructure:"SSL_MODE"` // For Postgres
	Charset    string `mapstructure:"CHARSET"`  // For MySQL
	BaseURL    string `mapstructure:"BASE_URL"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
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
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.Charset)
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
