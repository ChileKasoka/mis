package services

import (
	"fmt"
	"time"

	"github.com/ChileKasoka/mis/internal/models"
	"github.com/ChileKasoka/mis/internal/repositories"
	"github.com/google/uuid"
)

type VendorService interface {
	CreateVendor(vendor models.Vendor, userID uuid.UUID) (models.Vendor, error)
	GetVendorByID(id uuid.UUID) (models.Vendor, error)
	UpdateVendor(vendor models.Vendor) (models.Vendor, error)
	DeleteVendor(id uuid.UUID) error
}

type vendorServiceImpl struct {
	vendorRepo  repositories.VendorRepository
	userService UserService
}

func NewVendorService(vendorRepo repositories.VendorRepository, userService UserService) VendorService {
	return &vendorServiceImpl{
		vendorRepo:  vendorRepo,
		userService: userService,
	}
}

func (s *vendorServiceImpl) CreateVendor(vendor models.Vendor, userID uuid.UUID) (models.Vendor, error) {
	// Validate user ID
	if userID == uuid.Nil {
		return models.Vendor{}, fmt.Errorf("invalid user ID")
	}

	// Validate vendor fields
	// if vendor.Biography == "" {
	// 	return models.Vendor{}, fmt.Errorf("biography is required")
	// }

	// Check if the user exists
	_, err := s.userService.GetUser(userID.String())
	if err != nil {
		return models.Vendor{}, fmt.Errorf("user not found: %w", err)
	}

	// // Check if a vendor already exists for this user
	// existingVendor, err := s.vendorRepo.GetVendorByUserID(userID)
	// if err != nil && err.Error() != "vendor not found" {
	// 	// If the error is not "vendor not found," it means there's a database issue
	// 	return models.Vendor{}, fmt.Errorf("failed to check existing vendor: %w", err)
	// }
	// if existingVendor.ID != uuid.Nil {
	// 	// If a vendor already exists, return an error
	// 	return models.Vendor{}, fmt.Errorf("a vendor account already exists for this user")
	// }

	// Initialize vendor fields
	vendor.ID = uuid.New()
	vendor.CreatedAt = time.Now()
	vendor.UpdatedAt = time.Now()
	vendor.UserID = userID // Associate the vendor with the logged-in user

	// Create vendor in the repository
	createdVendor, err := s.vendorRepo.CreateVendor(vendor)
	if err != nil {
		return models.Vendor{}, fmt.Errorf("failed to create vendor: %w", err)
	}

	return createdVendor, nil
}

func (s *vendorServiceImpl) GetVendorByID(id uuid.UUID) (models.Vendor, error) {
	vendor, err := s.vendorRepo.GetVendorByID(id)
	if err != nil {
		return models.Vendor{}, fmt.Errorf("vendor not found: %w", err)
	}
	return vendor, nil
}

func (s *vendorServiceImpl) UpdateVendor(vendor models.Vendor) (models.Vendor, error) {
	vendor.UpdatedAt = time.Now()

	updatedVendor, err := s.vendorRepo.UpdateVendor(vendor)
	if err != nil {
		return models.Vendor{}, fmt.Errorf("failed to update vendor: %w", err)
	}
	return updatedVendor, nil
}

func (s *vendorServiceImpl) DeleteVendor(id uuid.UUID) error {
	err := s.vendorRepo.DeleteVendor(id)
	if err != nil {
		return fmt.Errorf("failed to delete vendor: %w", err)
	}
	return nil
}
