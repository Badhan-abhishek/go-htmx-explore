package lib

import (
	"net/http"

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
