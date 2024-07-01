package main

import (
	"github.com/nikhilryan/go-featuristic/api/routers"
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	dsn := config.GetDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.FeatureFlag{})
	if err != nil {
		return
	}

	cacheService := services.NewCacheService(cfg.CacheHost + ":" + cfg.CachePort)
	featureFlagService := services.NewFeatureFlagService(db, cacheService)

	r := routers.SetupRouter(featureFlagService)

	log.Println("Server running on port", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
