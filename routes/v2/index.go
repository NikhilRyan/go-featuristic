package v2

import (
	"github.com/go-chi/chi/v5"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
)

func Router(r chi.Router, featureFlagService *services.FeatureFlagService) {
	r.Route("/featureflags", func(r chi.Router) {
		FeatureFlagRouter(r, featureFlagService)
	})
}
