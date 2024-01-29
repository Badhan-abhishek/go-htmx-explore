package models

type Landlord struct {
	Model
	UserID     uint       `json:"userId,omitempty"`
	User       User       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Properties []Property `json:"properties"`
}

type Tenant struct {
	Model
	UserID          uint            `json:"userId,omitempty"`
	User            User            `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TenantAgreement TenantAgreement `json:"tenantAgreement"`
	TenantStatus    TenantStatus    `json:"tenantStatus"`
	TenantStatusID  uint            `json:"-"`
}

type User struct {
	Model
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Email       *string `json:"email" gorm:"unique"`
	Password    []byte  `json:"-"`
	PhoneNumber *string `json:"phoneNumber" gorm:"unique"`
	ProfilePic  *string
	Roles       []Role `json:"roles" gorm:"many2many:user_roles;"`
	Active      bool   `json:"active"`
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
