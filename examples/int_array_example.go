package examples

import (
	"encoding/json"
	"fmt"
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func RunIntArrayExample() {
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

	intArray := []int{1, 2, 3}
	intArrayJSON, err := json.Marshal(intArray)
	if err != nil {
		log.Fatalf("failed to marshal int array: %v", err)
	}
	intArrayFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "intArrayFeature",
		Value:     string(intArrayJSON),
		Type:      "intArray",
	}
	err = featureFlagService.CreateFlag(intArrayFlag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	value, err := featureFlagService.GetFlagValue("test", "intArrayFeature")
	if err != nil {
		log.Fatalf("failed to get feature flag value: %v", err)
	}
	fmt.Printf("Feature flag value: %v\n", value)
}
