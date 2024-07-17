package examples

import (
	"encoding/json"
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"log"
)

func RunIntArrayExample() {

	intArray := []int{1, 2, 3}
	intArrayJSON, err := json.Marshal(intArray)
	if err != nil {
		log.Fatalf("failed to marshal int array: %v", err)
	}
	intArrayFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "intArrayFeature",
		Value:     string(intArrayJSON),
		Type:      services.FlagTypeIntArray,
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
