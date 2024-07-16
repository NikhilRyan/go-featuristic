package tests

import (
	"bytes"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"github.com/nikhilryan/go-featuristic/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB
	featureFlagService *services.FeatureFlagService
	rolloutService     *services.RolloutService
)

func init() {
	dsn := "user=username password=password dbname=featureflag sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	cacheService := services.NewAppCacheService(client)
	featureFlagService = services.NewFeatureFlagService(db, cacheService)
	rolloutService = services.NewRolloutService(featureFlagService)
}

func setupRouter() *chi.Mux {
	return routes.InitializeRoutes(featureFlagService, rolloutService)
}

func TestCreateFlagAPI(t *testing.T) {
	router := setupRouter()

	flag := models.FeatureFlag{
		Namespace: "test",
		Key:       "feature1",
		Value:     "true",
		Type:      "bool",
	}
	flagJSON, _ := json.Marshal(flag)

	req, _ := http.NewRequest("POST", "/flags", bytes.NewBuffer(flagJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetFlagAPI(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/flags/test/feature1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetAllFlagsAPI(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/flags/test", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUpdateFlagAPI(t *testing.T) {
	router := setupRouter()

	flag := models.FeatureFlag{
		Namespace: "test",
		Key:       "feature1",
		Value:     "false",
		Type:      "bool",
	}
	flagJSON, _ := json.Marshal(flag)

	req, _ := http.NewRequest("PUT", "/flags/test/feature1", bytes.NewBuffer(flagJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteFlagAPI(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/flags/test/feature1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteAllFlagsAPI(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/flags/test", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetABTestVariantAPI(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/abtest/test/abTestFeature?user_id=user123&target_group=groupA", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestIsEnabledAPI(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/rollout/test/rolloutFeature?user_id=user123&rollout_percentage=50", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]bool
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if _, ok := response["enabled"]; !ok {
		t.Errorf("response does not contain enabled key")
	}
}
