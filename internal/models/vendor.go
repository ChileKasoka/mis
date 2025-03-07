package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Day string

const (
	Monday    Day = "Monday"
	Tuesday   Day = "Tuesday"
	Wednesday Day = "Wednesday"
	Thursday  Day = "Thursday"
	Friday    Day = "Friday"
	Saturday  Day = "Saturday"
	Sunday    Day = "Sunday"
)

type Vendor struct {
	ID             uuid.UUID      `json:"id"`
	UserID         uuid.UUID      `json:"user_id"`
	FirstName      string         `json:"first_name"` // Added from UI
	LastName       string         `json:"last_name"`  // Added from UI
	Email          string         `json:"email"`      // Added from UI
	Password       string         `json:"password"`
	Phone          string         `json:"phone"`              // Added from UI
	Address        string         `json:"address"`            // Added from UI
	Location       sql.NullString `json:"location,omitempty"` // Added from UI
	Website        sql.NullString `json:"website,omitempty"`  // Added from UI
	Biography      sql.NullString `json:"biography,omitempty"`
	ProfilePicture sql.NullString `json:"profile_picture,omitempty"`
	BusinessType   string         `json:"business_type"`            // Added from UI
	Experience     sql.NullString `json:"experience"`               // Added from UI
	Certification  sql.NullString `json:"certifications,omitempty"` // Added from UI
	Active         bool           `json:"active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

// NewVendor creates a new Vendor instance with the given parameters.
func NewVendor(
	userID uuid.UUID,
	firstName, lastName, email, password, phone, address, businessType string,
	location, website, biography, profilePicture, experience, certification sql.NullString,
) *Vendor {
	return &Vendor{
		ID:             uuid.New(),
		UserID:         userID,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Password:       password,
		Phone:          phone,
		Address:        address,
		Location:       location,
		Website:        website,
		Biography:      biography,
		ProfilePicture: profilePicture,
		BusinessType:   businessType,
		Experience:     experience,
		Certification:  certification,
		Active:         true, // Defaulting to active
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
