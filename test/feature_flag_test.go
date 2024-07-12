package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateFlag(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "feature1",
		Value:     "true",
		Type:      "bool",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"feature_flags\"").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.CreateFlag(flag)
	assert.NoError(t, err)
}

func TestGetFlag(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "feature1",
		Value:     "true",
		Type:      "bool",
	}

	rows := sqlmock.NewRows([]string{"namespace", "key", "value", "type"}).
		AddRow(flag.Namespace, flag.Key, flag.Value, flag.Type)
	mock.ExpectQuery("SELECT * FROM \"feature_flags\" WHERE").WillReturnRows(rows)

	result, err := service.GetFlag(flag.Namespace, flag.Key)
	assert.NoError(t, err)
	assert.Equal(t, flag, result)
}

func TestUpdateFlag(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "feature1",
		Value:     "false",
		Type:      "bool",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"feature_flags\" SET").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.UpdateFlag(flag)
	assert.NoError(t, err)
}

func TestDeleteFlag(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	namespace := "test"
	key := "feature1"

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"feature_flags\" WHERE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.DeleteFlag(namespace, key)
	assert.NoError(t, err)
}

func TestDeleteAllFlags(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	namespace := "test"

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"feature_flags\" WHERE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.DeleteAllFlags(namespace)
	assert.NoError(t, err)
}

func TestGetAllFlags(t *testing.T) {
	db, mock := setupTestDB(t)
	cache, _ := setupTestCache()
	service := services.NewFeatureFlagService(db, cache)

	flag := &models.FeatureFlag{
		Namespace: "test",
		Key:       "feature1",
		Value:     "true",
		Type:      "bool",
	}

	rows := sqlmock.NewRows([]string{"namespace", "key", "value", "type"}).
		AddRow(flag.Namespace, flag.Key, flag.Value, flag.Type)
	mock.ExpectQuery("SELECT * FROM \"feature_flags\" WHERE").WillReturnRows(rows)

	result, err := service.GetAllFlags(flag.Namespace)
	assert.NoError(t, err)
	assert.Equal(t, []*models.FeatureFlag{flag}, result)
}
