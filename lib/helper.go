package lib

import (
	"net/http"

	"com.backend/models"
	"github.com/gin-gonic/gin"
)

func GetAuthEmail(c *gin.Context) (*string, error) {
	currentEmail, exists := c.Get("Email")
	if !exists {
		err := NewHttpError("Something went wrong", "", http.StatusInternalServerError)
		return nil, err
	}
	email := currentEmail.(string)
	return &email, nil
}

func GetAuthLandlord(c *gin.Context) (*models.Landlord, error) {
	var user models.User
	var landlord models.Landlord
	currentEmail, err := GetAuthEmail(c)
	if err != nil {
		return nil, err
	}

	userResult := models.DB.Find(&user, "email = ?", currentEmail)
	if userResult.Error != nil {
		return nil, userResult.Error
	}

	landlordResult := models.DB.Find(&landlord, "user_id = ?", user.ID)
	if landlordResult.Error != nil {
		return nil, landlordResult.Error
	}

	return &landlord, nil
}
