package services

import (
	"encoding/json"
	"errors"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
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
	case FlagTypeInt:
		var intValue int
		if err := json.Unmarshal([]byte(flag.Value), &intValue); err != nil {
			return nil, err
		}
		value = intValue
	case FlagTypeFloat:
		var floatValue float64
		if err := json.Unmarshal([]byte(flag.Value), &floatValue); err != nil {
			return nil, err
		}
		value = floatValue
	case FlagTypeString:
		value = flag.Value
	case FlagTypeBool:
		var boolValue bool
		if err := json.Unmarshal([]byte(flag.Value), &boolValue); err != nil {
			return nil, err
		}
		value = boolValue
	case FlagTypeIntArray:
		var intArrayValue []int
		if err := json.Unmarshal([]byte(flag.Value), &intArrayValue); err != nil {
			return nil, err
		}
		value = intArrayValue
	case FlagTypeFloatArray:
		var floatArrayValue []float64
		if err := json.Unmarshal([]byte(flag.Value), &floatArrayValue); err != nil {
			return nil, err
		}
		value = floatArrayValue
	case FlagTypeStringArray:
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

func (s *FeatureFlagService) GetFunctionValue(namespace, key string, args ...interface{}) (interface{}, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return nil, err
	}

	if flag.Type != FlagTypeFunction {
		return flag.Value, nil
	}

	return CallFunction(flag.Value, args...)
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
	return s.cache.Delete(getCacheKey(namespace, key))
}

func (s *FeatureFlagService) DeleteAllFlags(namespace string) error {
	if err := s.db.Where("namespace = ?", namespace).Delete(&models.FeatureFlag{}).Error; err != nil {
		return err
	}

	// Invalidate cache for all flags under the namespace
	var flags []models.FeatureFlag
	if err := s.db.Where("namespace = ?", namespace).Find(&flags).Error; err == nil {
		for _, flag := range flags {
			err := s.cache.Delete(getCacheKey(namespace, flag.Key))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *FeatureFlagService) GetAllFlags(namespace string) ([]*models.FeatureFlag, error) {
	var flags []*models.FeatureFlag
	if err := s.db.Where("namespace = ?", namespace).Find(&flags).Error; err != nil {
		return nil, err
	}
	return flags, nil
}

func (s *FeatureFlagService) GetABTestVariant(namespace, key, userID, targetGroup string) (string, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return "", err
	}

	return determineABTestVariant(flag, userID, targetGroup)
}

func (s *FeatureFlagService) IsEnabled(namespace, key string, userID string, rolloutPercentage int) (bool, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return false, err
	}
	if flag.Type != "boolean" {
		return false, nil
	}
	hash := hashUserID(userID)
	return hash%100 < rolloutPercentage, nil
}

func (s *FeatureFlagService) GetRolloutPercentage(namespace, key string) (int, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	return 100, nil
}

func (s *FeatureFlagService) GetRolloutPercentageForUser(namespace, key, userID string) (int, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	hash := hashUserID(userID)
	return hash % 100, nil
}

func (s *FeatureFlagService) GetRolloutPercentageForUserAndNamespace(namespace, key, userID string) (int, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	hash := hashUserID(userID)
	return hash % 100, nil
}

func (s *FeatureFlagService) GetRolloutPercentageForUserAndKey(namespace, key, userID string) (int, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	hash := hashUserID(userID)
	return hash % 100, nil
}

func (s *FeatureFlagService) GetRolloutPercentageForUserAndNamespaceAndKey(namespace, key, userID string) (int, error) {
	flag, err := s.GetFlag(namespace, key)
	if err != nil {
		return 0, err
	}
	if flag.Type != "boolean" {
		return 0, nil
	}
	hash := hashUserID(userID)
	return hash % 100, nil
}
