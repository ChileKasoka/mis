package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Feedback struct {
	ID            uuid.UUID      `json:"id"`
	AppointmentID uuid.NullUUID  `json:"appointment_id"`
	Rating        sql.NullInt32  `json:"rating"`
	Comment       sql.NullString `json:"comment"`
}

type VendorAvailability struct {
	ID        uuid.UUID     `json:"id"`
	VendorID  uuid.NullUUID `json:"vendor_id"`
	DayOfWeek interface{}   `json:"day_of_week"`
	Date      sql.NullTime  `json:"date"`
}
