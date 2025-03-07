package services

import (
	"context"
	"errors"

	models "github.com/ChileKasoka/mis/internal/models"
	"github.com/ChileKasoka/mis/internal/repositories"
)

// AppointmentService defines the business logic layer for managing appointments
type AppointmentService interface {
	GetAllAppointments(ctx context.Context) ([]models.Appointment, error)
	GetAppointmentByID(id string) (*models.Appointment, error)
	CreateAppointment(appointment models.Appointment) (models.Appointment, error)
}

type appointmentServiceImpl struct {
	repository repositories.AppointmentRepository
}

// NewAppointmentService creates a new instance of AppointmentService
func NewAppointmentService(repository repositories.AppointmentRepository) AppointmentService {
	return &appointmentServiceImpl{repository: repository}
}

func (s *appointmentServiceImpl) CreateAppointment(appointment models.Appointment) (models.Appointment, error) {
	// Check if there's already a confirmed appointment for this vendor, date, and time slot
	exists, err := s.repository.CheckConfirmedAppointment(appointment.VendorID, appointment.TimeSlotID, appointment.Date)
	if err != nil {
		return models.Appointment{}, err
	}
	if exists {
		return models.Appointment{}, errors.New("an appointment is already confirmed for this vendor at the selected time")
	}

	// Proceed to create the appointment if no conflict exists
	createdAppointment, err := s.repository.CreateAppointment(appointment)
	if err != nil {
		return models.Appointment{}, err
	}
	return createdAppointment, nil
}

// GetAllAppointments retrieves all appointments from the repository
func (s *appointmentServiceImpl) GetAllAppointments(ctx context.Context) ([]models.Appointment, error) {
	appointments, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

// GetAppointmentByID retrieves a specific appointment by its ID
func (s *appointmentServiceImpl) GetAppointmentByID(id string) (*models.Appointment, error) {
	appointment, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}
