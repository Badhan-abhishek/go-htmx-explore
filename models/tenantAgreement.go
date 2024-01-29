package models

import "time"

type TenantAgreement struct {
	Model
	Label           string                `json:"label"`
	StartDate       time.Time             `json:"startDate"`
	EndDate         time.Time             `json:"endDate"`
	Status          TenantAgreementStatus `json:"status"`
	Lease           Document              `json:"lease"`
	LeaseID         uint                  `json:"-"`
	StatusID        uint                  `json:"-"`
	OtherDocuments  []Document            `json:"otherDocuments" gorm:"many2many:tenant_agreement_other_document;"`
	Rent            float32               `json:"rent"`
	TotalRentAmount float32               `json:"totalRentAmount"`
	DueDate         string                `json:"dueDate"`
	TenantID        uint                  `json:"tenantId"`
}

type TenantAgreementStatus struct {
	Model
	Name     string `json:"name"`
	Sequence int    `json:"sequence"`
}
