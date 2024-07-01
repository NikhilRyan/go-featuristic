package services

import (
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
	return s.cache.Set(flag.Namespace+"_"+flag.Key, flag.Value)
}

func (s *FeatureFlagService) GetFlag(namespace, key string) (*models.FeatureFlag, error) {
	var flag models.FeatureFlag
	if err := s.db.Where("namespace = ? AND key = ?", namespace, key).First(&flag).Error; err != nil {
		return nil, err
	}
	return &flag, nil
}

func (s *FeatureFlagService) UpdateFlag(flag *models.FeatureFlag) error {
	if err := s.db.Save(flag).Error; err != nil {
		return err
	}
	return s.cache.Set(flag.Namespace+"_"+flag.Key, flag.Value)
}

func (s *FeatureFlagService) DeleteFlag(namespace, key string) error {
	if err := s.db.Where("namespace = ? AND key = ?", namespace, key).Delete(&models.FeatureFlag{}).Error; err != nil {
		return err
	}
	return s.cache.Invalidate(namespace + "_" + key)
}
