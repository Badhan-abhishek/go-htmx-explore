package lib

import (
	"net/http"
	"strings"
	"time"

	"com.backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Payload struct {
	jwt.Claims
	Email         string    `json:"email"`
	Exp           time.Time `json:"exp"`
	Authenticated bool      `json:"authenticated"`
}

const signingKey = "asldfnlkajsndf"

func GetJwt(user *models.User) (string, error) {
	hmacSampleSecret := []byte(signingKey)
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
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	t := strings.Split(bearerToken, " ")
	token := t[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	result, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, NewHttpError("Invalid token", "", 401)
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"reason":  err.Error(),
		})
		c.Abort()
		return
	}
	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		// now := time.Now().Unix()
		// isTokenExpired := claims.VerifyExpiresAt(now, false)
		c.Set("Email", claims["email"])
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"reason":  "Invalid claims",
		})
		c.Abort()
	}
}
