package lib

// enums for document type
const (
	DocumentType_UnitImages = iota + 1
	DocumentType_Agreement
	DocumentType_Other
)

// enums for role
const (
	Role_Tenant = iota + 1
	Role_Landlord
	Role_Admin
)

// enums for tenant status
const (
	TenantStatus_Invited = iota + 1
	TenantStatus_Stale
	TenantStatus_InvitationAccepted
	TenantStatus_InTenantAgreement
	TenantStatus_Unknown
)
