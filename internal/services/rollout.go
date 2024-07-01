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
	_, err := h.Write([]byte(userID))
	if err != nil {
		return 0
	}
	return int(h.Sum32())
}

func (r *RolloutService) GetRolloutPercentage(namespace, key string) (int, error) {
	flag, err := r.featureFlagService.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	return 100, nil
}

func (r *RolloutService) GetRolloutPercentageForUser(namespace, key, userID string) (int, error) {
	flag, err := r.featureFlagService.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	hash := hashUserID(userID)
	return hash % 100, nil
}

func (r *RolloutService) GetRolloutPercentageForUserAndNamespace(namespace, key, userID string) (int, error) {
	flag, err := r.featureFlagService.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	hash := hashUserID(userID)
	return hash % 100, nil
}

func (r *RolloutService) GetRolloutPercentageForUserAndKey(namespace, key, userID string) (int, error) {
	flag, err := r.featureFlagService.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	hash := hashUserID(userID)
	return hash % 100, nil
}

func (r *RolloutService) GetRolloutPercentageForUserAndNamespaceAndKey(namespace, key, userID string) (int, error) {
	flag, err := r.featureFlagService.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	hash := hashUserID(userID)
	return hash % 100, nil
}
