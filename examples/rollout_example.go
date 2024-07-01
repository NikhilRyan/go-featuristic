package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func RunRolloutExample() {
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
	rolloutService := services.NewRolloutService(featureFlagService)

	// Create a new feature flag with a boolean value
	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "rolloutFeature",
		Value:     "true",
		Type:      "boolean",
	}
	err = featureFlagService.CreateFlag(flag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	// Simulate checking the feature flag status for multiple users
	rand.Seed(time.Now().UnixNano())
	userIDs := []string{"user1", "user2", "user3", "user4", "user5"}
	rolloutPercentage := 50

	for _, userID := range userIDs {
		enabled, err := rolloutService.IsEnabled("test", "rolloutFeature", userID, rolloutPercentage)
		if err != nil {
			log.Fatalf("failed to check rollout status: %v", err)
		}
		fmt.Printf("Feature flag status for %s: %v\n", userID, enabled)
	}
}
