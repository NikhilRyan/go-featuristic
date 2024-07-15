package examples

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func RunStringExample() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	dsn := config.GetDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.FeatureFlag{})
	if err != nil {
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr: cfg.CacheHost + ":" + cfg.CachePort,
	})
	cacheService := services.NewAppCacheService(client)
	featureFlagService := services.NewFeatureFlagService(db, cacheService)

	stringFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "stringFeature",
		Value:     "example string",
		Type:      services.FlagTypeString,
	}
	err = featureFlagService.CreateFlag(stringFlag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	value, err := featureFlagService.GetFlagValue("test", "stringFeature")
	if err != nil {
		log.Fatalf("failed to get feature flag value: %v", err)
	}
	fmt.Printf("Feature flag value: %v\n", value)
}
