package models

type Property struct {
	Model
	Label      string  `json:"label"`
	Address    Address `json:"address"`
	LandlordID uint    `json:"-"`
	Units      []Unit  `json:"units"`
}

type Address struct {
	Model
	Address1   string  `json:"address1"`
	Address2   *string `json:"address2"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postalCode"`
	State      string  `json:"state"`
	PropertyID uint    `json:"-"`
}

type Unit struct {
	Model
	Name        string     `json:"name"`
	Images      []Document `json:"images" gorm:"many2many:unit_image;"`
	PropertyID  uint       `json:"propertyId"`
	MaxCapacity int        `json:"maxCapacity"`
}
