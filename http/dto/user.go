package dto

type CreateUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	PhoneNumber *string `json:"phoneNumber"`
}

type SigninUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
