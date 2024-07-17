package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"log"
)

func RunFloatExample() {

	floatFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "floatFeature",
		Value:     "123.45",
		Type:      services.FlagTypeFloat,
	}
	err := featureFlagService.CreateFlag(floatFlag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	value, err := featureFlagService.GetFlagValue("test", "floatFeature")
	if err != nil {
		log.Fatalf("failed to get feature flag value: %v", err)
	}
	fmt.Printf("Feature flag value: %v\n", value)
}
