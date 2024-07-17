package services

import (
	"github.com/nikhilryan/go-featuristic/featuristic/models"
)

// FeatureFlagServiceProvider defines the interface for all feature flag services.
type FeatureFlagServiceProvider interface {
	// FeatureFlagService methods
	GetFlagValue(namespace, key string) (interface{}, error)
	GetFlag(namespace, key string) (*models.FeatureFlag, error)
	CreateFlag(flag *models.FeatureFlag) error
	UpdateFlag(flag *models.FeatureFlag) error
	DeleteFlag(namespace, key string) error
	GetAllFlags(namespace string) ([]*models.FeatureFlag, error)
	DeleteAllFlags(namespace string) error

	// RolloutService methods
	IsEnabled(namespace, key, userID string) (bool, error)
	GetRolloutPercentage(namespace, key string) (int, error)
	GetRolloutPercentageForUser(namespace, key, userID string) (int, error)
	GetRolloutPercentageForUserAndNamespace(namespace, key, userID string) (int, error)
	GetRolloutPercentageForUserAndKey(namespace, key, userID string) (int, error)
	GetRolloutPercentageForUserAndNamespaceAndKey(namespace, key, userID string) (int, error)

	// ABTestService methods
	GetABTestVariant(namespace, key, userID, targetGroup string) (string, error)
}
