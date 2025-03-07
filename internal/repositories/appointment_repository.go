package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sqlc "github.com/ChileKasoka/mis/db/sqlc"
	models "github.com/ChileKasoka/mis/internal/models"
	"github.com/google/uuid"
)

type AppointmentRepository interface {
	FindAll(ctx context.Context) ([]models.Appointment, error)
	FindById(id string) (*models.Appointment, error)
	CreateAppointment(appointment models.Appointment) (models.Appointment, error)
	CheckConfirmedAppointment(vendorID, timeSlotID uuid.UUID, date time.Time) (bool, error)
}

type appointmentRepositoryImpl struct {
	Queries *sqlc.Queries
}

func NewAppointmentRepository(db *sql.DB) AppointmentRepository {
	return &appointmentRepositoryImpl{Queries: sqlc.New(db)}
}

func (r *appointmentRepositoryImpl) CreateAppointment(appointment models.Appointment) (models.Appointment, error) {
	arg := sqlc.CreateAppointmentParams{
		ID:         uuid.New(),
		CustomerID: appointment.CustomerID,
		VendorID:   appointment.VendorID,
		Date:       appointment.Date,
		TimeSlotID: appointment.TimeSlotID,
		ServiceID:  appointment.ServiceID,
		Status:     string(appointment.Status),
	}

	createdAppointment, err := r.Queries.CreateAppointment(context.TODO(), arg)
	if err != nil {
		return models.Appointment{}, err
	}

	result := models.Appointment{
		ID:         createdAppointment.ID,
		CustomerID: createdAppointment.CustomerID,
		VendorID:   createdAppointment.VendorID,
		Date:       createdAppointment.Date,
		TimeSlotID: createdAppointment.TimeSlotID,
		ServiceID:  createdAppointment.ServiceID,
		Status:     models.AppointmentStatus(createdAppointment.Status),
	}

	return result, nil
}

func (r *appointmentRepositoryImpl) FindAll(ctx context.Context) ([]models.Appointment, error) {
	rows, err := r.Queries.GetAllAppointments(ctx)
	if err != nil {
		return nil, err
	}

	var appointments []models.Appointment
	for _, row := range rows {
		appointments = append(appointments, models.Appointment{
			ID:         row.ID,
			CustomerID: row.CustomerID,
			VendorID:   row.VendorID,
			Date:       row.Date,
			TimeSlotID: row.TimeSlotID,
			ServiceID:  row.ServiceID,
			Status:     models.AppointmentStatus(row.Status),
		})
	}

	return appointments, nil
}

func (r *appointmentRepositoryImpl) FindById(id string) (*models.Appointment, error) {
	// Convert the string ID to a UUID (assuming ID is stored as UUID)
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// Fetch the appointment using sqlc's generated method
	appointment, err := r.Queries.GetAppointmentById(context.TODO(), uuidID)
	if err != nil {
		return nil, err
	}

	// Map sqlc's Appointment result to your models.Appointment (if needed)
	result := &models.Appointment{
		ID:         appointment.ID,
		CustomerID: appointment.CustomerID,
		VendorID:   appointment.VendorID,
		Date:       appointment.Date,
		TimeSlotID: appointment.TimeSlotID,
		ServiceID:  appointment.ServiceID,
		Status:     models.AppointmentStatus(appointment.Status),
	}

	return result, nil
}

func (r *appointmentRepositoryImpl) CheckConfirmedAppointment(vendorID, timeSlotID uuid.UUID, date time.Time) (bool, error) {
	params := sqlc.CheckConfirmedAppointmentParams{
		VendorID:   vendorID,
		TimeSlotID: timeSlotID,
		Date:       date,
	}

	exists, err := r.Queries.CheckConfirmedAppointment(context.Background(), params)
	if err != nil {
		return false, fmt.Errorf("failed to check confirmed appointment: %w", err)
	}

	return exists, nil
}
