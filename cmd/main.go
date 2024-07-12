package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"github.com/nikhilryan/go-featuristic/routes"
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

	client := redis.NewClient(&redis.Options{
		Addr: cfg.CacheHost + ":" + cfg.CachePort,
	})
	cacheService := services.NewAppCacheService(client)

	featureFlagService := services.NewFeatureFlagService(db, cacheService)

	router := routes.InitializeRoutes(featureFlagService)

	log.Println("Server is running on port", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
