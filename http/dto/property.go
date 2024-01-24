package dto

type AddressInput struct {
	Address1   string  `json:"address1"`
	Address2   *string `json:"address2"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postalCode"`
	State      string  `json:"state"`
}

type CreatePropertyInput struct {
	Address AddressInput `json:"address"`
	Label   string       `json:"label"`
}

type UnitInput struct {
	Name string `json:"name"`
}
