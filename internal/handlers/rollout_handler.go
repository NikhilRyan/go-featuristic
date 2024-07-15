package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"net/http"
	"strconv"
)

type RolloutHandler struct {
	RolloutService *services.RolloutService
}

func NewRolloutHandler(rolloutService *services.RolloutService) *RolloutHandler {
	return &RolloutHandler{RolloutService: rolloutService}
}

func (h *RolloutHandler) IsEnabled(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	key := chi.URLParam(r, "key")
	userID := r.URL.Query().Get("user_id")
	rolloutPercentage, err := strconv.Atoi(r.URL.Query().Get("rollout_percentage"))
	if err != nil {
		http.Error(w, "Invalid rollout percentage", http.StatusBadRequest)
		return
	}

	enabled, err := h.RolloutService.IsEnabled(namespace, key, userID, rolloutPercentage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]bool{"enabled": enabled}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
