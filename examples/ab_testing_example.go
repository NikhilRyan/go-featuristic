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

func RunABExample() {
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

	// Create a new A/B test feature flag
	abTestFlag := &models.FeatureFlag{
		Namespace:    "test",
		Key:          "abTestFeature",
		Value:        "Variant A",
		ABTestValue:  "Variant B",
		ABTestType:   "A/B",
		TargetGroup:  "groupA",
		TargetGroupB: "groupB",
	}
	err = featureFlagService.CreateFlag(abTestFlag)
	if err != nil {
		log.Fatalf("failed to create A/B test feature flag: %v", err)
	}

	// Example users
	users := []struct {
		ID          string
		TargetGroup string
	}{
		{"user1", "groupA"},
		{"user2", "groupA"},
		{"user3", "groupB"},
		{"user4", "groupB"},
		{"user5", "groupA"},
	}

	// Determine the A/B test variant for each user
	for _, user := range users {
		variant, err := featureFlagService.GetABTestVariant("test", "abTestFeature", user.ID, user.TargetGroup)
		if err != nil {
			log.Printf("failed to get A/B test variant for %s: %v", user.ID, err)
			continue
		}
		fmt.Printf("A/B test variant for %s (group %s): %s\n", user.ID, user.TargetGroup, variant)
	}
}
