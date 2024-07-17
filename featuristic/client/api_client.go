package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"io"
	"net/http"
)

type FeatureFlagAPIClient struct {
	baseURL string
}

func NewFeatureFlagAPIClient(baseURL string) *FeatureFlagAPIClient {
	return &FeatureFlagAPIClient{
		baseURL: baseURL,
	}
}

// GetFlag retrieves a feature flag based on the namespace and key
func (c *FeatureFlagAPIClient) GetFlag(namespace, key string) (*models.FeatureFlag, error) {
	url := fmt.Sprintf("%s/flags/%s/%s", c.baseURL, namespace, key)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get flag: %v", resp.Status)
	}

	var flag models.FeatureFlag
	if err := json.NewDecoder(resp.Body).Decode(&flag); err != nil {
		return nil, err
	}
	return &flag, nil
}

// CreateFlag creates a new feature flag
func (c *FeatureFlagAPIClient) CreateFlag(flag *models.FeatureFlag) error {
	url := fmt.Sprintf("%s/flags", c.baseURL)
	flagJSON, err := json.Marshal(flag)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(flagJSON))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create flag: %v", resp.Status)
	}

	return nil
}

// UpdateFlag updates an existing feature flag
func (c *FeatureFlagAPIClient) UpdateFlag(namespace, key string, flag *models.FeatureFlag) error {
	url := fmt.Sprintf("%s/flags/%s/%s", c.baseURL, namespace, key)
	flagJSON, err := json.Marshal(flag)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(flagJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update flag: %v", resp.Status)
	}

	return nil
}

// DeleteFlag deletes a feature flag
func (c *FeatureFlagAPIClient) DeleteFlag(namespace, key string) error {
	url := fmt.Sprintf("%s/flags/%s/%s", c.baseURL, namespace, key)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete flag: %v", resp.Status)
	}

	return nil
}

// GetAllFlags retrieves all feature flags in a namespace
func (c *FeatureFlagAPIClient) GetAllFlags(namespace string) ([]*models.FeatureFlag, error) {
	url := fmt.Sprintf("%s/flags/%s", c.baseURL, namespace)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get all flags: %v", resp.Status)
	}

	var flags []*models.FeatureFlag
	if err := json.NewDecoder(resp.Body).Decode(&flags); err != nil {
		return nil, err
	}
	return flags, nil
}

// DeleteAllFlags deletes all feature flags in a namespace
func (c *FeatureFlagAPIClient) DeleteAllFlags(namespace string) error {
	url := fmt.Sprintf("%s/flags/%s", c.baseURL, namespace)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete all flags: %v", resp.Status)
	}

	return nil
}

// IsRolloutEnabled checks if a rollout is enabled for a given namespace, key, and user ID
func (c *FeatureFlagAPIClient) IsRolloutEnabled(namespace, key, userID string) (bool, error) {
	url := fmt.Sprintf("%s/rollout/%s/%s?user_id=%s", c.baseURL, namespace, key, userID)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to check rollout: %v", resp.Status)
	}

	var result struct {
		Enabled bool `json:"enabled"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}
	return result.Enabled, nil
}

// GetABTestVariant retrieves the A/B test variant for a given namespace, key, and user ID
func (c *FeatureFlagAPIClient) GetABTestVariant(namespace, key, userID string) (string, error) {
	url := fmt.Sprintf("%s/abtest/%s/%s?user_id=%s", c.baseURL, namespace, key, userID)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get A/B test variant: %v", resp.Status)
	}

	var result struct {
		Variant string `json:"variant"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Variant, nil
}
