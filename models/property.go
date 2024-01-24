package models

type Property struct {
	Model
	Label      string  `json:"label"`
	Address    Address `json:"address" gorm:"foreignKey:ID;references:ID"`
	LandlordID uint    `json:"landlordId"`
}

type Address struct {
	Model
	Address1   string  `json:"address1"`
	Address2   *string `json:"address2"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postalCode"`
	State      string  `json:"state"`
}

type Unit struct {
	Model
	Name     string     `json:"name"`
	Property Property   `json:"property" gorm:"foreignKey:ID;references:ID"`
	Images   []Document `json:"images" gorm:"many2many:unit_image;"`
}
