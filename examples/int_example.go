package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"log"
)

func RunIntExample() {

	intFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "intFeature",
		Value:     "123",
		Type:      services.FlagTypeInt,
	}
	err := featureFlagService.CreateFlag(intFlag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	value, err := featureFlagService.GetFlagValue("test", "intFeature")
	if err != nil {
		log.Fatalf("failed to get feature flag value: %v", err)
	}
	fmt.Printf("Feature flag value: %v\n", value)
}
