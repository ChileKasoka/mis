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

type VendorRepository interface {
	CreateVendor(vendor models.Vendor) (models.Vendor, error)
	GetVendorByID(id uuid.UUID) (models.Vendor, error)
	UpdateVendor(vendor models.Vendor) (models.Vendor, error)
	DeleteVendor(id uuid.UUID) error
}

type vendorRepositoryImpl struct {
	Queries *sqlc.Queries
}

func NewVendorRepository(db *sql.DB) VendorRepository {
	return &vendorRepositoryImpl{
		Queries: sqlc.New(db),
	}
}

func convertDaysToStrings(days []models.Day) []string {
	strDays := make([]string, len(days))
	for i, day := range days {
		strDays[i] = string(day)
	}
	return strDays
}

func convertStringsToDays(strDays []string) []models.Day {
	days := make([]models.Day, len(strDays))
	for i, strDay := range strDays {
		days[i] = models.Day(strDay)
	}
	return days
}

func (r *vendorRepositoryImpl) CreateVendor(vendor models.Vendor) (models.Vendor, error) {
	arg := sqlc.CreateVendorParams{
		ID:             uuid.New(),
		FirstName:      vendor.FirstName,
		LastName:       vendor.LastName,
		Email:          vendor.Email,
		Password:       vendor.Password,
		Phone:          vendor.Phone,
		Address:        vendor.Address,
		Location:       vendor.Location,
		Website:        vendor.Website,
		Biography:      vendor.Biography,
		ProfilePicture: vendor.ProfilePicture,
		BusinessType:   vendor.BusinessType,
		Experience:     vendor.Experience,
		Certification:  vendor.Certification,
		Active:         vendor.Active,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	createdVendor, err := r.Queries.CreateVendor(context.TODO(), arg)
	if err != nil {
		return models.Vendor{}, fmt.Errorf("error creating vendor: %w", err)
	}

	return models.Vendor{
		ID:             createdVendor.ID,
		FirstName:      createdVendor.FirstName,
		LastName:       createdVendor.LastName,
		Email:          createdVendor.Email,
		Biography:      createdVendor.Biography,
		ProfilePicture: createdVendor.ProfilePicture,
		Active:         createdVendor.Active,
		CreatedAt:      createdVendor.CreatedAt,
		UpdatedAt:      createdVendor.UpdatedAt,
	}, nil
}

func (r *vendorRepositoryImpl) GetVendorByID(id uuid.UUID) (models.Vendor, error) {
	vendor, err := r.Queries.GetVendor(context.TODO(), id)
	if err != nil {
		return models.Vendor{}, fmt.Errorf("error retrieving vendor with ID %v: %w", id, err)
	}

	return models.Vendor{
		ID:             vendor.ID,
		FirstName:      vendor.FirstName,
		LastName:       vendor.LastName,
		Email:          vendor.Email,
		Biography:      vendor.Biography,
		ProfilePicture: vendor.ProfilePicture,
		Active:         vendor.Active,
		CreatedAt:      vendor.CreatedAt,
		UpdatedAt:      vendor.UpdatedAt,
	}, nil
}

func (r *vendorRepositoryImpl) UpdateVendor(vendor models.Vendor) (models.Vendor, error) {
	arg := sqlc.UpdateVendorParams{
		ID:             vendor.ID,
		Biography:      vendor.Biography,
		ProfilePicture: vendor.ProfilePicture,
		Active:         vendor.Active,
		UpdatedAt:      time.Now(),
	}

	updatedVendor, err := r.Queries.UpdateVendor(context.TODO(), arg)
	if err != nil {
		return models.Vendor{}, fmt.Errorf("error updating vendor with ID %v: %w", vendor.ID, err)
	}

	return models.Vendor{
		ID:             updatedVendor.ID,
		Biography:      updatedVendor.Biography,
		ProfilePicture: updatedVendor.ProfilePicture,
		Active:         updatedVendor.Active,
		UpdatedAt:      updatedVendor.UpdatedAt,
	}, nil
}

func (r *vendorRepositoryImpl) DeleteVendor(id uuid.UUID) error {
	_, err := r.Queries.DeleteVendor(context.TODO(), id)
	if err != nil {
		return fmt.Errorf("error deleting vendor with ID %v: %w", id, err)
	}
	return nil
}
