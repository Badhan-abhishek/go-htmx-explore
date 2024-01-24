package models

import "time"

type TenantAgreement struct {
	Model
	Label           string                `json:"label"`
	StartDate       time.Time             `json:"startDate"`
	EndDate         time.Time             `json:"endDate"`
	Status          TenantAgreementStatus `json:"status" gorm:"foreignKey:ID;references:ID"`
	Lease           Document              `json:"lease" gorm:"foreignKey:ID;references:ID"`
	OtherDocuments  []Document            `json:"otherDocuments" gorm:"many2many:tenant_agreement_other_document;"`
	Rent            float32               `json:"rent"`
	TotalRentAmount float32               `json:"totalRentAmount"`
	DueDate         string                `json:"dueDate"`
}

type TenantAgreementStatus struct {
	Model
	Name     string `json:"name"`
	Sequence int    `json:"sequence" gorm:"unique"`
}
