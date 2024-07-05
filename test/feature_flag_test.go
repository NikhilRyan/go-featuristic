package tests

import (
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateFlag(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	err = db.AutoMigrate(&models.FeatureFlag{})
	if err != nil {
		return
	}

	cache := services.NewCacheService("localhost:6379")
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{Namespace: "test", Key: "new_feature", Value: "true", Type: "boolean"}
	err = service.CreateFlag(flag)
	if err != nil {
		t.Fatalf("failed to create feature flag: %v", err)
	}

	retrievedFlag, err := service.GetFlag("test", "new_feature")
	if err != nil {
		t.Fatalf("failed to retrieve feature flag: %v", err)
	}
	if retrievedFlag.Value != "true" {
		t.Errorf("expected flag value to be 'true', got '%s'", retrievedFlag.Value)
	}
}

func TestUpdateFlag(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	err = db.AutoMigrate(&models.FeatureFlag{})
	if err != nil {
		return
	}

	cache := services.NewCacheService("localhost:6379")
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{Namespace: "test", Key: "new_feature", Value: "true", Type: "boolean"}
	err = service.CreateFlag(flag)
	if err != nil {
		return
	}

	flag.Value = "false"
	err = service.UpdateFlag(flag)
	if err != nil {
		t.Fatalf("failed to update feature flag: %v", err)
	}

	retrievedFlag, err := service.GetFlag("test", "new_feature")
	if err != nil {
		t.Fatalf("failed to retrieve feature flag: %v", err)
	}
	if retrievedFlag.Value != "false" {
		t.Errorf("expected flag value to be 'false', got '%s'", retrievedFlag.Value)
	}
}

func TestDeleteFlag(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	err = db.AutoMigrate(&models.FeatureFlag{})
	if err != nil {
		return
	}

	cache := services.NewCacheService("localhost:6379")
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{Namespace: "test", Key: "new_feature", Value: "true", Type: "boolean"}
	err = service.CreateFlag(flag)
	if err != nil {
		return
	}

	err = service.DeleteFlag("test", "new_feature")
	if err != nil {
		t.Fatalf("failed to delete feature flag: %v", err)
	}

	_, err = service.GetFlag("test", "new_feature")
	if err == nil {
		t.Fatalf("expected error when retrieving deleted feature flag, got none")
	}
}
