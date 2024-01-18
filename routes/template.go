package routes

import (
	"com.quizApp/http/views"
	"github.com/gin-gonic/gin"
)

func InitializeTemplateRoute(r *gin.Engine) {
	r.GET("/", views.Home)
	r.GET("/clicked", views.Swap)
	r.POST("sign-in", views.SignIn)
}
