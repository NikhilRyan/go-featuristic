package middleware

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"net/http"
	"strconv"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	err := validate.RegisterValidation("flagtype", validateFlagType)
	if err != nil {
		return
	}
}

func ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

			if err := validate.Struct(flag); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := validateFlagValue(flag); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func validateFlagType(fl validator.FieldLevel) bool {
	flagType := fl.Field().String()
	switch flagType {
	case "int", "float", "string", "bool", "intArray", "floatArray", "stringArray":
		return true
	}
	return false
}

func validateFlagValue(flag models.FeatureFlag) error {
	switch flag.Type {
	case "int":
		_, err := strconv.Atoi(flag.Value)
		if err != nil {
			return err
		}
	case "float":
		_, err := strconv.ParseFloat(flag.Value, 64)
		if err != nil {
			return err
		}
	case "string":
		// No validation needed for string
	case "bool":
		_, err := strconv.ParseBool(flag.Value)
		if err != nil {
			return err
		}
	case "intArray":
		var intArray []int
		err := json.Unmarshal([]byte(flag.Value), &intArray)
		if err != nil {
			return err
		}
	case "floatArray":
		var floatArray []float64
		err := json.Unmarshal([]byte(flag.Value), &floatArray)
		if err != nil {
			return err
		}
	case "stringArray":
		var stringArray []string
		err := json.Unmarshal([]byte(flag.Value), &stringArray)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported flag type")
	}
	return nil
}
