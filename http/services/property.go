package services

import (
	"com.backend/http/dto"
)

func AddPropertyService(input dto.CreatePropertyInput) {
	// addressInput := input.Address
	// address := models.Address{Address1: addressInput.Address1,
	// 	Address2: addressInput.Address2, City: addressInput.City,
	// 	Country: addressInput.Country, PostalCode: addressInput.PostalCode,
	// 	State: addressInput.State}
	// err := models.DB.Transaction(func(tx *gorm.DB) error {
	// 	addressResult := tx.Create(&address)
	// 	if addressResult.Error != nil {
	// 		return addressResult.Error
	// 	}

	// 	addressID := address.ID

	// 	property := models.Property{Label: input.Label, Address: address}

	// 	return nil
	// })
}
