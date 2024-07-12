package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestGetABTestVariant(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{
		Namespace:   "test",
		Key:         "abTestFeature",
		Value:       "Variant A",
		ABTestValue: "Variant B",
		ABTestType:  "A/B",
		Type:        "string",
	}

	rows := sqlmock.NewRows([]string{"namespace", "key", "value", "type", "ab_test_value", "ab_test_type"}).
		AddRow(flag.Namespace, flag.Key, flag.Value, flag.Type, flag.ABTestValue, flag.ABTestType)
	mock.ExpectQuery("SELECT * FROM \"feature_flags\" WHERE").WillReturnRows(rows)

	userID := "user123"
	targetGroup := "groupA"
	variant, err := service.GetABTestVariant(flag.Namespace, flag.Key, userID, targetGroup)
	assert.NoError(t, err)
	assert.Contains(t, []string{flag.Value, flag.ABTestValue}, variant)
}

func TestGetABTestVariantWithDifferentUsers(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{
		Namespace:   "test",
		Key:         "abTestFeature",
		Value:       "Variant A",
		ABTestValue: "Variant B",
		ABTestType:  "A/B",
		Type:        "string",
	}

	rows := sqlmock.NewRows([]string{"namespace", "key", "value", "type", "ab_test_value", "ab_test_type"}).
		AddRow(flag.Namespace, flag.Key, flag.Value, flag.Type, flag.ABTestValue, flag.ABTestType)
	mock.ExpectQuery("SELECT * FROM \"feature_flags\" WHERE").WillReturnRows(rows)

	users := []string{"user123", "user456", "user789"}
	targetGroup := "groupA"
	for _, userID := range users {
		variant, err := service.GetABTestVariant(flag.Namespace, flag.Key, userID, targetGroup)
		assert.NoError(t, err)
		assert.Contains(t, []string{flag.Value, flag.ABTestValue}, variant)
	}
}

func TestGetABTestVariantWithNonABTestFlag(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "nonABTestFeature",
		Value:     "Feature Value",
		Type:      "string",
	}

	rows := sqlmock.NewRows([]string{"namespace", "key", "value", "type"}).
		AddRow(flag.Namespace, flag.Key, flag.Value, flag.Type)
	mock.ExpectQuery("SELECT * FROM \"feature_flags\" WHERE").WillReturnRows(rows)

	userID := "user123"
	targetGroup := "groupA"
	variant, err := service.GetABTestVariant(flag.Namespace, flag.Key, userID, targetGroup)
	assert.Error(t, err)
	assert.Equal(t, "", variant)
}
