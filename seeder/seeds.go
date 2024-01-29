package main

import (
	"log"

	"com.backend/lib"
	"com.backend/models"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func SeedAll() []Seed {
	return []Seed{
		{
			Name: "Document Types",
			Run: func(db *gorm.DB) error {
				documentTypes := []models.DocumentType{
					{
						Name:     "Unit Image",
						Sequence: lib.DocumentType_UnitImages,
					},
					{
						Name:     "Agreement",
						Sequence: lib.DocumentType_Agreement,
					},
					{
						Name:     "Other",
						Sequence: lib.DocumentType_Other,
					},
				}
				result := db.Create(&documentTypes)
				if result.Error != nil {
					return result.Error
				}
				return nil
			},
		},
		{
			Name: "Roles",
			Run: func(db *gorm.DB) error {
				roles := []models.Role{
					{
						Name:     "Tenant",
						Sequence: lib.Role_Tenant,
					},
					{
						Name:     "Landlord",
						Sequence: lib.Role_Landlord,
					},
					{
						Name:     "Admin",
						Sequence: lib.Role_Admin,
					},
				}
				if result := db.Create(&roles); result.Error != nil {
					return result.Error
				}
				return nil
			},
		},
		{
			Name: "Tenant Statuses",
			Run: func(db *gorm.DB) error {
				tenantStatuses := []models.TenantStatus{
					{
						Name:     "In Tenant Agreement",
						Sequence: lib.TenantStatus_InTenantAgreement,
					},
					{
						Name:     "Invited",
						Sequence: lib.TenantStatus_Invited,
					},
					{
						Name:     "Invitation Accepted",
						Sequence: lib.TenantStatus_InvitationAccepted,
					},
					{
						Name:     "Stale",
						Sequence: lib.TenantStatus_Stale,
					},
					{
						Name:     "Unknown",
						Sequence: lib.TenantStatus_Unknown,
					},
				}
				if result := db.Create(&tenantStatuses); result.Error != nil {
					return result.Error
				}
				return nil
			},
		},
	}
}

func main() {
	models.ConnectDatabase()
	db := models.DB
	for _, seed := range SeedAll() {
		if err := seed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	DB, _ := db.DB()
	DB.Close()
}
