package main

import (
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/config/cache"
	"github.com/nikhilryan/go-featuristic/config/db"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"github.com/nikhilryan/go-featuristic/routes"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	database := db.GetDB()
	redisClient := cache.GetRedisClient()

	cacheService := services.NewAppCacheService(redisClient)
	featureFlagService := services.NewFeatureFlagService(database, cacheService)
	rolloutService := services.NewRolloutService(featureFlagService)

	router := routes.InitializeRoutes(featureFlagService, rolloutService)

	log.Println("Server is running on port", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
