package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ChileKasoka/mis/internal/models"
	"github.com/ChileKasoka/mis/internal/services"
	"github.com/ChileKasoka/mis/middleware"
	"github.com/ChileKasoka/mis/util"

	"github.com/go-chi/chi"
)

type VendorHandler struct {
	VendorService services.VendorService
	JWTSecret     string
}

func NewVendorHandler(service services.VendorService, jwtSecret string) *VendorHandler {
	return &VendorHandler{
		VendorService: service,
		JWTSecret:     jwtSecret,
	}
}

func (v *VendorHandler) RegisterRoutes(r chi.Router) {
	r.With(middleware.JWTAuth(v.JWTSecret)).Post("/new-vendor", v.HandleNewVendor)
	r.With(middleware.JWTAuth(v.JWTSecret)).Get("/vendor/{id}", v.GetCurrentVendor)
	r.With(middleware.JWTAuth(v.JWTSecret)).Put("/update-vendor", v.UpdateVendorDetails)
}

func (v *VendorHandler) HandleNewVendor(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from context
	user, err := util.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Decode request body
	var vendor models.Vendor
	if err := json.NewDecoder(r.Body).Decode(&vendor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service layer
	createdVendor, err := v.VendorService.CreateVendor(vendor, user)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdVendor)
}

func (v *VendorHandler) GetCurrentVendor(w http.ResponseWriter, r *http.Request) {

}

func (v *VendorHandler) UpdateVendorDetails(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized: missing or invalid user ID", http.StatusUnauthorized)
		return
	}
}
