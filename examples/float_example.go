package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func RunFloatExample() {
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

	cacheService := services.NewCacheService(cfg.CacheHost + ":" + cfg.CachePort)
	featureFlagService := services.NewFeatureFlagService(db, cacheService)

	floatFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "floatFeature",
		Value:     "123.45",
		Type:      "float",
	}
	err = featureFlagService.CreateFlag(floatFlag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	value, err := featureFlagService.GetFlagValue("test", "floatFeature")
	if err != nil {
		log.Fatalf("failed to get feature flag value: %v", err)
	}
	fmt.Printf("Feature flag value: %v\n", value)
}
