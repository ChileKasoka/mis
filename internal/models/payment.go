package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethod string

type PaymentStatus string

const (
	MobileMoney PaymentMethod = "mobile_money"
	Visa        PaymentMethod = "visa"
	Offline     PaymentMethod = "offline"
)

const (
	pending  PaymentStatus = "pending"
	paid     PaymentStatus = "paid"
	reversed PaymentStatus = "reversed"
)

type Payment struct {
	ID            uuid.UUID     `json:"id"`
	AppointmentID uuid.UUID     `json:"appointment_id"`
	CustomerID    uuid.UUID     `json:"customer_id"`
	VendorID      uuid.UUID     `json:"vendor_id"`
	Amount        string        `json:"amount"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	Status        PaymentStatus `json:"status"`
	Date          time.Time     `json:"date"`
	TransactionID time.Time     `json:"transaction_id"`
}

func NewPayment(appointmentId uuid.UUID, customerId uuid.UUID, vendorId uuid.UUID, amount, paymentMethod PaymentMethod, status PaymentStatus, date time.Time, transactionId string) *Payment {
	if paymentMethod == "" {
		paymentMethod = Offline
	}

	if status == "" {
		status = pending
	}
	return &Payment{}
}
