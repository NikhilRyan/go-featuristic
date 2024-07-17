package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"log"
)

func RunStringExample() {

	stringFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "stringFeature",
		Value:     "example string",
		Type:      services.FlagTypeString,
	}
	err := featureFlagService.CreateFlag(stringFlag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	value, err := featureFlagService.GetFlagValue("test", "stringFeature")
	if err != nil {
		log.Fatalf("failed to get feature flag value: %v", err)
	}
	fmt.Printf("Feature flag value: %v\n", value)
}
