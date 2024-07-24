package routes

import (
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"github.com/nikhilryan/go-featuristic/handlers"
	"github.com/nikhilryan/go-featuristic/middleware"
)

func InitializeRoutes(router Router, featureFlagService *services.FeatureFlagService) {
	// Apply middleware
	router.Use(middleware.ValidationMiddleware)

	// Create handler
	handler := handlers.NewFeatureFlagHandler(featureFlagService)

	// Define routes
	router.Get("/flags/{namespace}/{key}", handler.GetFlag)
	router.Post("/flags", handler.CreateFlag)
	router.Get("/flags/{namespace}", handler.GetAllFlags)
	router.Put("/flags/{namespace}/{key}", handler.UpdateFlag)
	router.Delete("/flags/{namespace}/{key}", handler.DeleteFlag)
	router.Delete("/flags/{namespace}", handler.DeleteAllFlags)
	router.Get("/abtest/{namespace}/{key}", handler.GetABTestVariant)
	router.Get("/rollout/{namespace}/{key}", handler.IsEnabled)
}
