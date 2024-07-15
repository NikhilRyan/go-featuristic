package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nikhilryan/go-featuristic/internal/handlers"
	"github.com/nikhilryan/go-featuristic/internal/middleware"
	"github.com/nikhilryan/go-featuristic/internal/services"
)

func InitializeRoutes(featureFlagService *services.FeatureFlagService, rolloutService *services.RolloutService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.ValidationMiddleware) // Add the validation middleware

	handler := handlers.NewFeatureFlagHandler(featureFlagService)
	rolloutHandler := handlers.NewRolloutHandler(rolloutService)

	// Feature flag routes
	r.Post("/flags", handler.CreateFlag)
	r.Get("/flags/{namespace}/{key}", handler.GetFlag)
	r.Get("/flags/{namespace}", handler.GetAllFlags)
	r.Put("/flags/{namespace}/{key}", handler.UpdateFlag)
	r.Delete("/flags/{namespace}/{key}", handler.DeleteFlag)
	r.Delete("/flags/{namespace}", handler.DeleteAllFlags)

	// A/B testing routes
	r.Get("/abtest/{namespace}/{key}", handler.GetABTestVariant)

	// Rollout routes
	r.Get("/rollout/{namespace}/{key}", rolloutHandler.IsEnabled)

	return r
}
