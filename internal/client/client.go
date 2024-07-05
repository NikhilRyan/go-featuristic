package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"io"
	"net/http"
)

type FeatureFlagClient struct {
	baseURL string
}

func NewFeatureFlagClient(baseURL string) *FeatureFlagClient {
	return &FeatureFlagClient{baseURL: baseURL}
}

func (c *FeatureFlagClient) GetFlag(namespace, key string) (*models.FeatureFlag, error) {
	resp, err := http.Get(c.baseURL + "/flags?namespace=" + namespace + "&key=" + key)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	var flag models.FeatureFlag
	if err := json.NewDecoder(resp.Body).Decode(&flag); err != nil {
		return nil, err
	}
	return &flag, nil
}

func (c *FeatureFlagClient) CreateFlag(flag *models.FeatureFlag) error {
	body, err := json.Marshal(flag)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.baseURL+"/flags", "application/json", bytes.NewBuffer(body))
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
		return errors.New("failed to create feature flag")
	}

	return nil
}

func (c *FeatureFlagClient) UpdateFlag(flag *models.FeatureFlag) error {
	body, err := json.Marshal(flag)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, c.baseURL+"/flags", bytes.NewBuffer(body))
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
		return errors.New("failed to update feature flag")
	}

	return nil
}

func (c *FeatureFlagClient) DeleteFlag(namespace, key string) error {
	req, err := http.NewRequest(http.MethodDelete, c.baseURL+"/flags?namespace="+namespace+"&key="+key, nil)
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
		return errors.New("failed to delete feature flag")
	}

	return nil
}
