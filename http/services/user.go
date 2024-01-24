package services

import (
	"fmt"

	"com.backend/http/dto"
	"com.backend/lib"
	"com.backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateUserServiceResponse struct {
	user  *models.User
	token string
}

type UserWithLandlord struct {
	models.User
	Landlord models.Landlord `json:"landlord"`
}

func CreateUserService(input dto.CreateUserInput) (*CreateUserServiceResponse, error) {
	password := input.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return nil, lib.NewHttpError("Password Hashing Failed", err.Error(), 500)
	}
	user := models.User{FirstName: input.FirstName, LastName: input.LastName, Email: &input.Email, Password: hash}
	landlord := models.Landlord{}

	err = models.DB.Transaction(func(tx *gorm.DB) error {
		userResult := tx.Create(&user)
		if userResult.Error != nil {
			return userResult.Error
		}

		landlord.UserID = user.ID

		landlordResult := tx.Create(&landlord)

		if landlordResult.Error != nil {
			return landlordResult.Error
		}

		return nil
	})

	if err != nil {
		return nil, lib.NewHttpError("User Creation Failed", err.Error(), 500)
	}

	token, tokenGenErr := lib.GetJwt(&user)

	if tokenGenErr != nil {
		return nil, lib.NewHttpError("Token Generation Failed", tokenGenErr.Error(), 500)
	}

	return &CreateUserServiceResponse{
		user:  &user,
		token: token,
	}, nil
}

func SignInUserService(input dto.SigninUserInput) (*string, error) {
	var user models.User
	result := models.DB.Find(&user, "email = ?", input.Email)
	if result.Error != nil {
		return nil, lib.NewHttpError("Invalid Credentials", result.Error.Error(), 401)
	}

	passCheckErr := bcrypt.CompareHashAndPassword(user.Password, []byte(input.Password))
	if passCheckErr != nil {
		return nil, lib.NewHttpError("Invalid Credentials", "", 401)
	}

	token, tokenGenErr := lib.GetJwt(&user)

	if tokenGenErr != nil {
		return nil, lib.NewHttpError("Token Generation Failed", tokenGenErr.Error(), 500)
	}

	return &token, nil
}

func GetAllUsersService() ([]models.User, error) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		return nil, lib.NewHttpError("User Fetch Failed", result.Error.Error(), 500)
	}
	return users, nil
}

func GetUserService(email string) (*models.User, error) {
	var user models.User
	result := models.DB.Find(&user, "email = ?", email)
	if result.Error != nil {
		return nil, lib.NewHttpError("Failed to get user", result.Error.Error(), 404)
	}
	if user.Email == nil {
		return nil, lib.NewHttpError("User not found", "", 404)
	}
	return &user, nil
}

func UpdateUserService(email *string, input dto.UpdateUserInput) (*models.User, error) {
	var user models.User
	result := models.DB.Find(&user, "email = ?", email)
	if result.Error != nil {
		return nil, lib.NewHttpError("Failed to get user", result.Error.Error(), 404)
	}

	if input.PhoneNumber != nil {
		user.PhoneNumber = *input.PhoneNumber
	}

	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		user.LastName = *input.LastName
	}

	result = models.DB.Save(&user)
	if result.Error != nil {
		return nil, lib.NewHttpError("Failed to update user", result.Error.Error(), 500)
	}

	return &user, nil
}

func GetUserWithLandlordService(email string) (*UserWithLandlord, error) {
	var userWithLandlord UserWithLandlord

	err := models.DB.Where("email = ?", email).First(&userWithLandlord.User).Error

	if err != nil {
		return nil, lib.NewHttpError("Failed to get user", err.Error(), 404)
	}

	err = models.DB.Debug().Omit("user, user_id").Where("user_id = ?", userWithLandlord.User.ID).First(&userWithLandlord.Landlord).Error

	if err != nil {
		// Do not throw error if not landlord
		fmt.Printf("err: %v\n", err)
	}

	if userWithLandlord.Email == nil {
		return nil, lib.NewHttpError("User not found", "", 404)
	}
	return &userWithLandlord, nil
}
