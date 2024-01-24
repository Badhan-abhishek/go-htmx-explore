package middleware

import (
	"com.backend/lib"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		lib.VerifyJwt(c)
	}
}
