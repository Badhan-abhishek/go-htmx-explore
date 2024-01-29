package dto

type AddressInput struct {
	Address1   string  `json:"address1" binding:"required"`
	Address2   *string `json:"address2"`
	City       string  `json:"city" binding:"required"`
	Country    string  `json:"country" binding:"required"`
	PostalCode string  `json:"postalCode" binding:"required"`
	State      string  `json:"state" binding:"required"`
}

type CreatePropertyInput struct {
	Address AddressInput `json:"address" binding:"required"`
	Label   string       `json:"label" binding:"required"`
}

type UnitInput struct {
	Name        string     `json:"name" binding:"required"`
	PropertyID  uint       `json:"propertyId" binding:"required"`
	MaxCapacity int        `json:"maxCapacity"`
	Images      []Document `json:"images"`
}

type Document struct {
	TypeID     uint    `json:"typeId"`
	Url        *string `json:"url"`
	Name       string  `json:"name"`
	ExternalID *string `json:"externalId"`
}
