package models

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewContact(userID, phone, address string) *Contact {
	return &Contact{
		ID:        uuid.New(),
		UserID:    uuid.MustParse(userID),
		Phone:     phone,
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
