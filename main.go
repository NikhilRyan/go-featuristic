package main

import (
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikhilryan/go-featuristic/config"
	"github.com/nikhilryan/go-featuristic/config/db"
	"github.com/nikhilryan/go-featuristic/featuristic/client"
	"github.com/nikhilryan/go-featuristic/routes/v2"
)

func main() {
	cfg, err := config.LoadConfig("config/")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	database := db.GetDB()
	// Initialize Redis UniversalClient
	redisOptions := &redis.UniversalOptions{
		Addrs: []string{"localhost:6379"},
	}
	redisClient := redis.NewUniversalClient(redisOptions)
	cacheService := services.NewAppCacheService(services.NewRedisUniversalClientAdapter(redisClient))

	featureFlagService := services.NewFeatureFlagService(database, cacheService)
	funcClient := client.NewFeatureFlagFuncClient(featureFlagService)
	selectedClient := funcClient

	r := chi.NewRouter()
	r.Route("/v2", func(r chi.Router) {
		v2.Router(r, selectedClient.FeatureFlagService)
	})

	log.Println("Server is running on port", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
