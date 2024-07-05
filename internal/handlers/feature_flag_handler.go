package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"github.com/nikhilryan/go-featuristic/internal/services"
	"net/http"
)

type FeatureFlagHandler struct {
	FeatureFlagService *services.FeatureFlagService
}

func NewFeatureFlagHandler(featureFlagService *services.FeatureFlagService) *FeatureFlagHandler {
	return &FeatureFlagHandler{FeatureFlagService: featureFlagService}
}

func (h *FeatureFlagHandler) CreateFlag(w http.ResponseWriter, r *http.Request) {
	var flag models.FeatureFlag
	if err := json.NewDecoder(r.Body).Decode(&flag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.FeatureFlagService.CreateFlag(&flag); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *FeatureFlagHandler) GetFlag(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	key := chi.URLParam(r, "key")
	flag, err := h.FeatureFlagService.GetFlag(namespace, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(flag)
	if err != nil {
		return
	}
}

func (h *FeatureFlagHandler) GetAllFlags(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	flags, err := h.FeatureFlagService.GetAllFlags(namespace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(flags)
	if err != nil {
		return
	}
}

func (h *FeatureFlagHandler) UpdateFlag(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	key := chi.URLParam(r, "key")
	var flag models.FeatureFlag
	if err := json.NewDecoder(r.Body).Decode(&flag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	flag.Namespace = namespace
	flag.Key = key
	if err := h.FeatureFlagService.UpdateFlag(&flag); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *FeatureFlagHandler) DeleteFlag(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	key := chi.URLParam(r, "key")
	if err := h.FeatureFlagService.DeleteFlag(namespace, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *FeatureFlagHandler) DeleteAllFlags(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	if err := h.FeatureFlagService.DeleteAllFlags(namespace); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *FeatureFlagHandler) GetABTestVariant(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	key := chi.URLParam(r, "key")
	userID := r.URL.Query().Get("user_id")
	targetGroup := r.URL.Query().Get("target_group")
	variant, err := h.FeatureFlagService.GetABTestVariant(namespace, key, userID, targetGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(map[string]string{"variant": variant})
	if err != nil {
		return
	}
}
