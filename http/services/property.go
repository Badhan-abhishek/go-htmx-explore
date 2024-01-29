package services

import (
	"net/http"

	"com.backend/http/dto"
	"com.backend/lib"
	"com.backend/models"
	"gorm.io/gorm"
)

func AddPropertyService(input dto.CreatePropertyInput, landlord *models.Landlord) (*models.Property, error) {
	addressInput := input.Address
	address := models.Address{Address1: addressInput.Address1,
		Address2: addressInput.Address2, City: addressInput.City,
		Country: addressInput.Country, PostalCode: addressInput.PostalCode,
		State: addressInput.State}

	property := models.Property{Label: input.Label, LandlordID: landlord.ID}
	err := models.DB.Transaction(func(tx *gorm.DB) error {

		propertyResult := tx.Create(&property)
		if propertyResult.Error != nil {
			return propertyResult.Error
		}

		address.PropertyID = property.ID
		addressResult := tx.Create(&address)
		if addressResult.Error != nil {
			return addressResult.Error
		}
		return nil
	})

	if err != nil {
		return nil, lib.NewHttpError("Failed to create property", err.Error(), http.StatusBadRequest)
	}

	// Merge address into property
	property.Address = address

	return &property, nil
}

func GetLandlordPropertiesService(landlord *models.Landlord) (*[]models.Property, error) {
	var properties []models.Property

	propertiesResult := models.DB.Preload("Address").Preload("Units").Find(&properties, "landlord_id = ?", landlord.ID)
	if propertiesResult.Error != nil {
		return nil, lib.NewHttpError("Failed to fetch properties", propertiesResult.Error.Error(), http.StatusBadRequest)
	}
	return &properties, nil
}

func GetPropertyByID(propertyId uint) (*models.Property, error) {
	var property models.Property
	propertyResult := models.DB.Find(&property, "ID = ?", propertyId)
	if propertyResult.Error != nil {
		return nil, lib.NewHttpError("Cannot find property", propertyResult.Error.Error(), http.StatusBadRequest)
	}
	return &property, nil
}

func AddUnitToPropertyService(input dto.UnitInput) (*models.Unit, error) {
	_, err := GetPropertyByID(input.PropertyID)

	if err != nil {
		err = lib.NewHttpError("Cannot find property", err.Error(), http.StatusBadRequest)
		return nil, err
	}

	unit := models.Unit{Name: input.Name, PropertyID: input.PropertyID, MaxCapacity: input.MaxCapacity}
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		if unitResult := tx.Create(&unit); unitResult.Error != nil {
			return unitResult.Error
		}
		for _, doc := range input.Images {
			newDoc := models.Document{Url: doc.Url, TypeID: doc.TypeID, Name: &doc.Name}
			if docResult := tx.Create(&newDoc); docResult.Error != nil {
				return docResult.Error
			}
			tx.Model(&unit).Association("Images").Append(&newDoc)
		}
		return nil
	})

	if err != nil {
		return nil, lib.NewHttpError("Failed to create unit", err.Error(), http.StatusBadRequest)
	}

	return &unit, nil
}
