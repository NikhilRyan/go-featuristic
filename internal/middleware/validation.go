package middleware

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"

	"github.com/nikhilryan/go-featuristic/internal/models"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Basic validation logic
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
				return
			}

			var flag models.FeatureFlag
			if err := json.NewDecoder(r.Body).Decode(&flag); err != nil {
				http.Error(w, "Invalid JSON format", http.StatusBadRequest)
				return
			}

			// Validate the feature flag model
			if err := validate.Struct(flag); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		// Validate the presence of required parameters for GET requests
		if r.Method == http.MethodGet {
			if strings.Contains(r.URL.Path, "/flags/") {
				namespace := r.URL.Query().Get("namespace")
				key := r.URL.Query().Get("key")
				if namespace == "" || key == "" {
					http.Error(w, "Namespace and key are required", http.StatusBadRequest)
					return
				}
			}
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
