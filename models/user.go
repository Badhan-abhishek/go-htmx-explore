package models

type Landlord struct {
	Model
	UserID     uint       `json:"user_id,omitempty"`
	User       User       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Properties []Property `json:"properties"`
}

type Tenant struct {
	Model
	UserID          uint            `json:"user_id,omitempty"`
	User            User            `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TenantAgreement TenantAgreement `json:"tenantAgreement" gorm:"foreignKey:ID;references:ID"`
	TenantStatus    TenantStatus    `json:"tenantStatus" gorm:"foreignKey:ID;references:ID"`
}

type User struct {
	Model
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Email       *string `json:"email" gorm:"unique"`
	Password    []byte  `json:"-"`
	PhoneNumber string  `json:"phoneNumber" gor:"unique"`
	ProfilePic  string  `json:"profilePic"`
	Roles       []Role  `json:"roles" gorm:"many2many:user_roles;"`
	Active      bool    `json:"active"`
}

type Role struct {
	Model
	Name     string `json:"name"`
	Sequence int    `json:"sequence"`
}
type TenantStatus struct {
	Model
	Name     string `json:"name"`
	Sequence int    `json:"sequence"`
}
