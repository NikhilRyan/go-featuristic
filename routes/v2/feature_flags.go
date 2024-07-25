package v2

import (
	"github.com/go-chi/chi/v5"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"github.com/nikhilryan/go-featuristic/handlers"
	"github.com/nikhilryan/go-featuristic/middleware"
)

func FeatureFlagRouter(r chi.Router, featureFlagService *services.FeatureFlagService) {
	handler := handlers.NewFeatureFlagHandler(featureFlagService)

	r.With(middleware.ValidationMiddleware).Route("/", func(r chi.Router) {
		r.Post("/", handler.CreateFlag)
		r.Get("/{namespace}/{key}", handler.GetFlag)
		r.Get("/{namespace}", handler.GetAllFlags)
		r.Put("/{namespace}/{key}", handler.UpdateFlag)
		r.Delete("/{namespace}/{key}", handler.DeleteFlag)
		r.Delete("/{namespace}", handler.DeleteAllFlags)
	})

	r.Get("/abtest/{namespace}/{key}", handler.GetABTestVariant)
	r.Get("/rollout/{namespace}/{key}", handler.IsEnabled)
}
