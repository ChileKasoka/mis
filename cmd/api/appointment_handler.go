package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	//sqlc "github.com/ChileKasoka/mis/db/sqlc"
	"github.com/ChileKasoka/mis/internal/services"
	"github.com/ChileKasoka/mis/util"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type AppointmentHandler struct {
	appointmentService services.AppointmentService
}

func NewAppointmentHandler(appointmentService services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{appointmentService: appointmentService}
}

func (h *AppointmentHandler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.GetAllAppointments) // GET /appointments
}

func (h *AppointmentHandler) GetAllAppointments(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	appointments, err := h.appointmentService.GetAllAppointments(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

// GetAppointmentById handles the request to get an appointment by its ID
func (h *AppointmentHandler) GetAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Parse the appointment ID from the URL
	appointmentID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("could not parse id")
		return
	}

	// Convert the UUID to a string before passing it to the service
	appointment, err := h.appointmentService.GetAppointmentByID(appointmentID.String())
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "could not decode")
		return
	}

	// Return the appointment as a JSON response
	util.RespondWithJSON(w, http.StatusOK, appointment)
}
