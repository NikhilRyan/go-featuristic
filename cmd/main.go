package main

import (
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/config/db"
	"github.com/nikhilryan/go-featuristic/featuristic/client"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"github.com/nikhilryan/go-featuristic/routes"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize database
	database := db.GetDB()

	// Initialize Redis UniversalClient
	redisOptions := &redis.UniversalOptions{
		Addrs: []string{"localhost:6379"},
	}
	redisClient := redis.NewUniversalClient(redisOptions)
	cacheService := services.NewAppCacheService(services.NewRedisUniversalClientAdapter(redisClient))

	// Alternatively, for redis.Client:
	// redisClient := redis.NewClient(&redis.Options{
	//     Addr: "localhost:6379",
	// })
	// cacheService := services.NewAppCacheService(services.NewRedisClientAdapter(redisClient))

	// Initialize FeatureFlagService with the adapter
	featureFlagService := services.NewFeatureFlagService(database, cacheService)

	// Use function call client
	funcClient := client.NewFeatureFlagFuncClient(featureFlagService)
	// Use API client
	// apiClient := client.NewFeatureFlagAPIClient(cfg.BaseURL)

	// Choose the client to use
	selectedClient := funcClient // or apiClient

	// Initialize Chi router and routes
	chiRouter := routes.NewChiRouter()
	routes.InitializeRoutes(chiRouter, selectedClient.FeatureFlagService)

	log.Println("Server is running on port", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, chiRouter))
}
