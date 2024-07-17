package client

import (
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
)

type FeatureFlagFuncClient struct {
	FeatureFlagService *services.FeatureFlagService
	RolloutService     *services.RolloutService
}

func NewFeatureFlagFuncClient(featureFlagService *services.FeatureFlagService, rolloutService *services.RolloutService) *FeatureFlagFuncClient {
	return &FeatureFlagFuncClient{
		FeatureFlagService: featureFlagService,
		RolloutService:     rolloutService,
	}
}

// GetFlag retrieves a feature flag based on the namespace and key
func (c *FeatureFlagFuncClient) GetFlag(namespace, key string) (*models.FeatureFlag, error) {
	return c.FeatureFlagService.GetFlag(namespace, key)
}

// CreateFlag creates a new feature flag
func (c *FeatureFlagFuncClient) CreateFlag(flag *models.FeatureFlag) error {
	return c.FeatureFlagService.CreateFlag(flag)
}

// UpdateFlag updates an existing feature flag
func (c *FeatureFlagFuncClient) UpdateFlag(flag *models.FeatureFlag) error {
	return c.FeatureFlagService.UpdateFlag(flag)
}

// DeleteFlag deletes a feature flag
func (c *FeatureFlagFuncClient) DeleteFlag(namespace, key string) error {
	return c.FeatureFlagService.DeleteFlag(namespace, key)
}

// GetAllFlags retrieves all feature flags in a namespace
func (c *FeatureFlagFuncClient) GetAllFlags(namespace string) ([]models.FeatureFlag, error) {
	return c.FeatureFlagService.GetAllFlags(namespace)
}

// DeleteAllFlags deletes all feature flags in a namespace
func (c *FeatureFlagFuncClient) DeleteAllFlags(namespace string) error {
	return c.FeatureFlagService.DeleteAllFlags(namespace)
}

// IsRolloutEnabled checks if a rollout is enabled for a given namespace, key, and user ID
func (c *FeatureFlagFuncClient) IsRolloutEnabled(namespace, key, userID string, rolloutPercentage int) (bool, error) {
	return c.RolloutService.IsEnabled(namespace, key, userID, rolloutPercentage)
}

// GetABTestVariant retrieves the A/B test variant for a given namespace, key, and user ID
func (c *FeatureFlagFuncClient) GetABTestVariant(namespace, key, userID, targetGroup string) (string, error) {
	return c.FeatureFlagService.GetABTestVariant(namespace, key, userID, targetGroup)
}
