package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"gorm.io/gorm"
)

type FeatureFlagService struct {
	db    *gorm.DB
	cache *CacheService
}

func NewFeatureFlagService(db *gorm.DB, cache *CacheService) *FeatureFlagService {
	return &FeatureFlagService{db: db, cache: cache}
}

func (s *FeatureFlagService) CreateFlag(flag *models.FeatureFlag) error {
	if err := s.db.Create(flag).Error; err != nil {
		return err
	}
	flagJSON, err := json.Marshal(flag)
	if err != nil {
		return err
	}
	return s.cache.Set(getCacheKey(flag.Namespace, flag.Key), string(flagJSON))
}

func (s *FeatureFlagService) GetFlag(namespace, key string) (*models.FeatureFlag, error) {
	cacheKey := getCacheKey(namespace, key)

	// Check cache first
	flagJSON, err := s.cache.Get(cacheKey)
	if err == nil {
		var flag models.FeatureFlag
		if err := json.Unmarshal([]byte(flagJSON), &flag); err == nil {
			return &flag, nil
		}
	}

	// If not in cache, fetch from DB
	var flag models.FeatureFlag
	if err := s.db.Where("namespace = ? AND key = ?", namespace, key).First(&flag).Error; err != nil {
		return nil, err
	}

	// Update cache
	flagJSONBytes, err := json.Marshal(&flag)
	if err == nil {
		_ = s.cache.Set(cacheKey, string(flagJSONBytes))
	}

	return &flag, nil
}

func (s *FeatureFlagService) GetFlagValue(namespace, key string) (interface{}, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return nil, err
	}

	var value interface{}
	switch flag.Type {
	case "int":
		var intValue int
		if err := json.Unmarshal([]byte(flag.Value), &intValue); err != nil {
			return nil, err
		}
		value = intValue
	case "float":
		var floatValue float64
		if err := json.Unmarshal([]byte(flag.Value), &floatValue); err != nil {
			return nil, err
		}
		value = floatValue
	case "string":
		value = flag.Value
	case "intArray":
		var intArrayValue []int
		if err := json.Unmarshal([]byte(flag.Value), &intArrayValue); err != nil {
			return nil, err
		}
		value = intArrayValue
	case "floatArray":
		var floatArrayValue []float64
		if err := json.Unmarshal([]byte(flag.Value), &floatArrayValue); err != nil {
			return nil, err
		}
		value = floatArrayValue
	case "stringArray":
		var stringArrayValue []string
		if err := json.Unmarshal([]byte(flag.Value), &stringArrayValue); err != nil {
			return nil, err
		}
		value = stringArrayValue
	default:
		return nil, errors.New("unsupported type")
	}

	return value, nil
}

func (s *FeatureFlagService) UpdateFlag(flag *models.FeatureFlag) error {
	if err := s.db.Save(flag).Error; err != nil {
		return err
	}
	flagJSON, err := json.Marshal(flag)
	if err != nil {
		return err
	}
	return s.cache.Set(getCacheKey(flag.Namespace, flag.Key), string(flagJSON))
}

func (s *FeatureFlagService) DeleteFlag(namespace, key string) error {
	if err := s.db.Where("namespace = ? AND key = ?", namespace, key).Delete(&models.FeatureFlag{}).Error; err != nil {
		return err
	}
	return s.cache.Invalidate(getCacheKey(namespace, key))
}

func (s *FeatureFlagService) GetAllFlags(namespace string) ([]models.FeatureFlag, error) {
	var flags []models.FeatureFlag
	if err := s.db.Where("namespace = ?", namespace).Find(&flags).Error; err != nil {
		return nil, err
	}
	return flags, nil
}

func (s *FeatureFlagService) GetFlagCount(namespace string) (int64, error) {
	var count int64
	if err := s.db.Model(&models.FeatureFlag{}).Where("namespace = ?", namespace).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (s *FeatureFlagService) GetFlagKeys(namespace string) ([]string, error) {
	var keys []string
	if err := s.db.Model(&models.FeatureFlag{}).Where("namespace = ?", namespace).Pluck("key", &keys).Error; err != nil {
		return nil, err
	}
	return keys, nil
}

func (s *FeatureFlagService) GetFlagKeysByType(namespace string, flagType string) ([]string, error) {
	var keys []string
	if err := s.db.Model(&models.FeatureFlag{}).Where("namespace = ? AND type = ?", namespace, flagType).Pluck("key", &keys).Error; err != nil {
		return nil, err
	}
	return keys, nil
}

func (s *FeatureFlagService) GetFlagKeysByTypeAndValue(namespace string, flagType string, value interface{}) ([]string, error) {
	var keys []string
	if err := s.db.Model(&models.FeatureFlag{}).Where("namespace = ? AND type = ? AND value = ?", namespace, flagType, value).Pluck("key", &keys).Error; err != nil {
		return nil, err
	}
	return keys, nil
}

func getCacheKey(namespace, key string) string {
	return fmt.Sprintf("%s_%s", namespace, key)
}
