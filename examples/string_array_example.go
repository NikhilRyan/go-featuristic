package examples

import (
	"encoding/json"
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"log"
)

func RunStringArrayExample() {

	stringArray := []string{"feature1", "feature2", "feature3"}
	stringArrayJSON, err := json.Marshal(stringArray)
	if err != nil {
		log.Fatalf("failed to marshal string array: %v", err)
	}
	stringArrayFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "stringArrayFeature",
		Value:     string(stringArrayJSON),
		Type:      services.FlagTypeStringArray,
	}
	err = featureFlagService.CreateFlag(stringArrayFlag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	value, err := featureFlagService.GetFlagValue("test", "stringArrayFeature")
	if err != nil {
		log.Fatalf("failed to get feature flag value: %v", err)
	}
	fmt.Printf("Feature flag value: %v\n", value)
}
