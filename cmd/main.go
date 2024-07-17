package main

import (
	"github.com/nikhilryan/go-featuristic/config/cache"
	"github.com/nikhilryan/go-featuristic/config/db"
	"log"
	"net/http"

	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/featuristic/client"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"github.com/nikhilryan/go-featuristic/routes"
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

	// Use function call client
	funcClient := client.NewFeatureFlagFuncClient(featureFlagService)
	// Use API client
	//apiClient := client.NewFeatureFlagAPIClient(cfg.BaseURL)

	// Choose the client to use
	selectedClient := funcClient // or apiClient

	router := routes.InitializeRoutes(selectedClient.FeatureFlagService)

	log.Println("Server is running on port", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
