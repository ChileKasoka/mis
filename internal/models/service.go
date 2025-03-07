package models

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID          uuid.UUID `json:"id"`
	VendorID    uuid.UUID `json:"vendor_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewService(vendorId uuid.UUID, name string, description string, duration int, price float64) *Service {

	// Default duration to 60 minutes (1 hour) if not provided
	if duration == 0 {
		duration = 60
	}

	// Default price to 0.00 if not provided
	if price == 0 {
		price = 0.00
	}

	return &Service{
		ID:          uuid.New(),
		VendorID:    vendorId,
		Name:        name,
		Description: description,
		Price:       price,
		Duration:    duration,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
