package examples

import (
	"encoding/json"
	"fmt"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"log"
)

func RunFloatArrayExample() {

	floatArray := []float64{1.1, 2.2, 3.3}
	floatArrayJSON, err := json.Marshal(floatArray)
	if err != nil {
		log.Fatalf("failed to marshal float array: %v", err)
	}
	floatArrayFlag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "floatArrayFeature",
		Value:     string(floatArrayJSON),
		Type:      services.FlagTypeFloatArray,
	}
	err = featureFlagService.CreateFlag(floatArrayFlag)
	if err != nil {
		log.Fatalf("failed to create feature flag: %v", err)
	}

	value, err := featureFlagService.GetFlagValue("test", "floatArrayFeature")
	if err != nil {
		log.Fatalf("failed to get feature flag value: %v", err)
	}
	fmt.Printf("Feature flag value: %v\n", value)
}
