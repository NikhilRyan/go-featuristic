package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"log"
	"math/rand"
	"time"
)

func RunRolloutExample() {

	// Create a new feature flag with a boolean value
	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "rolloutFeature",
		Value:     "true",
		Type:      "boolean",
	}
	err := featureFlagService.CreateFlag(flag)
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
