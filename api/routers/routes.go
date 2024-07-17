package routers

import (
	"github.com/gorilla/mux"
	"github.com/nikhilryan/go-featuristic/api/handlers"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
)

func SetupRouter(featureFlagService *services.FeatureFlagService) *mux.Router {
	featureFlagHandler := handlers.NewFeatureFlagHandler(featureFlagService)

	r := mux.NewRouter()
	r.HandleFunc("/flags", featureFlagHandler.CreateFlag).Methods("POST")
	r.HandleFunc("/flags", featureFlagHandler.GetFlag).Methods("GET")
	r.HandleFunc("/flags", featureFlagHandler.UpdateFlag).Methods("PUT")
	r.HandleFunc("/flags", featureFlagHandler.DeleteFlag).Methods("DELETE")

	return r
}
