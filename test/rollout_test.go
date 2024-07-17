package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"github.com/stretchr/testify/assert"
)

func TestIsEnabled(t *testing.T) {
	db, mock := setupTestDB(t)
	cacheService, _ := setupTestCache()
	featureFlagService := services.NewFeatureFlagService(db, cacheService)

	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "rolloutFeature",
		Value:     "true",
		Type:      "boolean",
	}

	rows := sqlmock.NewRows([]string{"namespace", "key", "value", "type"}).
		AddRow(flag.Namespace, flag.Key, flag.Value, flag.Type)
	mock.ExpectQuery("SELECT * FROM \"feature_flags\" WHERE").WillReturnRows(rows)

	userID := "user123"
	rolloutPercentage := 50
	enabled, err := featureFlagService.IsEnabled(flag.Namespace, flag.Key, userID, rolloutPercentage)
	assert.NoError(t, err)
	assert.IsType(t, bool(enabled), enabled)
}

func TestIsEnabledWithDifferentRolloutPercentages(t *testing.T) {
	db, mock := setupTestDB(t)
	cacheService, _ := setupTestCache()
	featureFlagService := services.NewFeatureFlagService(db, cacheService)

	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "rolloutFeature",
		Value:     "true",
		Type:      "boolean",
	}

	rows := sqlmock.NewRows([]string{"namespace", "key", "value", "type"}).
		AddRow(flag.Namespace, flag.Key, flag.Value, flag.Type)
	mock.ExpectQuery("SELECT * FROM \"feature_flags\" WHERE").WillReturnRows(rows)

	userID := "user123"

	// Test with 0% rollout
	enabled, err := featureFlagService.IsEnabled(flag.Namespace, flag.Key, userID, 0)
	assert.NoError(t, err)
	assert.Equal(t, false, enabled)

	// Test with 100% rollout
	enabled, err = featureFlagService.IsEnabled(flag.Namespace, flag.Key, userID, 100)
	assert.NoError(t, err)
	assert.Equal(t, true, enabled)
}

func TestIsEnabledWithNonBooleanFlag(t *testing.T) {
	db, mock := setupTestDB(t)
	cacheService, _ := setupTestCache()
	featureFlagService := services.NewFeatureFlagService(db, cacheService)

	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "rolloutFeature",
		Value:     "42",
		Type:      "int",
	}

	rows := sqlmock.NewRows([]string{"namespace", "key", "value", "type"}).
		AddRow(flag.Namespace, flag.Key, flag.Value, flag.Type)
	mock.ExpectQuery("SELECT * FROM \"feature_flags\" WHERE").WillReturnRows(rows)

	userID := "user123"
	rolloutPercentage := 50
	enabled, err := featureFlagService.IsEnabled(flag.Namespace, flag.Key, userID, rolloutPercentage)
	assert.Error(t, err)
	assert.Equal(t, false, enabled)
	assert.EqualError(t, err, "feature flag is not a boolean type")
}
