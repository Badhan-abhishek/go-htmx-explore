package lib

import (
	"fmt"
	"net/http"
	"time"

	"com.quizApp/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Payload struct {
	jwt.Claims
	email         string    `json:"email"`
	exp           time.Time `json:"exp"`
	authenticated bool      `json:"authenticated"`
}

func GetJwt(user *models.User) (string, error) {
	hmacSampleSecret := []byte("asldfnlkajsndf")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":         user.Email,
		"exp":           time.Now().UTC().Add(time.Hour * 48).Unix(),
		"authenticated": true,
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJwt(c *gin.Context) {
	token := c.Request.Header.Get("token")
	fmt.Printf("%T", token)
	if token != "" {
		token, err := jwt.ParseWithClaims(token, &Payload{}, func(*jwt.Token) (interface{}, error) {
			return []byte("All your base"), nil
		})
		if claims, ok := token.Claims.(*Payload); ok && token.Valid {
			fmt.Printf("%v %v", claims.email, claims.exp)
		} else {
			fmt.Println(err)
		}
		c.Next()
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
	return
}
