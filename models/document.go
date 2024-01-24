package models

type Document struct {
	Model
	Url        *string      `json:"url"`
	Type       DocumentType `json:"type" gorm:"foreignKey:ID;references:ID"`
	Name       *string      `json:"name"`
	ExternalID *string      `json:"externalId"`
}

type DocumentType struct {
	Model
	Name string `json:"name"`
}
