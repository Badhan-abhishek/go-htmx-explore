package services

import (
	"com.quizApp/http/dto"
	"com.quizApp/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUserService(c *gin.Context, input dto.CreateUserInput) (*gorm.DB, error) {
	password := input.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return nil, err
	}

	user := models.User{Name: input.Name, Email: &input.Email, Password: hash}
	result := models.DB.Create(&user)
	return result, nil
}
