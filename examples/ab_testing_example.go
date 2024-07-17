package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"log"
)

func RunABExample() {

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
	err := featureFlagService.CreateFlag(abTestFlag)
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
