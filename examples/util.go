package examples

import (
	"github.com/nikhilryan/go-featuristic/config/cache"
	"github.com/nikhilryan/go-featuristic/config/db"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
)

var (
	featureFlagService *services.FeatureFlagService
	rolloutService     *services.RolloutService
)

func init() {

	database := db.GetDB()
	redisClient := cache.GetRedisClient()

	cacheService := services.NewAppCacheService(redisClient)
	featureFlagService = services.NewFeatureFlagService(database, cacheService)
	rolloutService = services.NewRolloutService(featureFlagService)
}
