package routes

import (
	"github.com/go-chi/chi"
	"github.com/nikhilryan/go-featuristic/internal/handlers"
	"github.com/nikhilryan/go-featuristic/internal/middleware"
	"github.com/nikhilryan/go-featuristic/internal/services"
)

func InitializeRoutes(featureFlagService *services.FeatureFlagService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.ValidationMiddleware) // Add the validation middleware

	handler := handlers.NewFeatureFlagHandler(featureFlagService)

	r.Post("/flags", handler.CreateFlag)
	r.Get("/flags/{namespace}/{key}", handler.GetFlag)
	r.Get("/flags/{namespace}", handler.GetAllFlags)
	r.Put("/flags/{namespace}/{key}", handler.UpdateFlag)
	r.Delete("/flags/{namespace}/{key}", handler.DeleteFlag)
	r.Delete("/flags/{namespace}", handler.DeleteAllFlags)
	r.Get("/abtest/{namespace}/{key}", handler.GetABTestVariant)

	return r
}
