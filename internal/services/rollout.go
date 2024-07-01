package services

import (
	"hash/fnv"
)

type RolloutService struct {
	featureFlagService *FeatureFlagService
}

func NewRolloutService(featureFlagService *FeatureFlagService) *RolloutService {
	return &RolloutService{featureFlagService: featureFlagService}
}

func (r *RolloutService) IsEnabled(namespace, key string, userID string, rolloutPercentage int) (bool, error) {
	flag, err := r.featureFlagService.GetFlag(namespace, key)
	if err != nil {
		return false, err
	}
	if flag.Type != "boolean" {
		return false, nil
	}
	hash := hashUserID(userID)
	return hash%100 < rolloutPercentage, nil
}

func hashUserID(userID string) int {
	h := fnv.New32a()
	h.Write([]byte(userID))
	return int(h.Sum32())
}
