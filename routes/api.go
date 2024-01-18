package routes

import (
	"com.quizApp/http/api/controller"
	"github.com/gin-gonic/gin"
)

func InitializeAPIRoute(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/sign-in", controller.Signin)
		api.POST("/sign-up", controller.Signup)
		api.GET("/get-users", controller.GetAllUsers)
		api.GET("/get-users/:email", controller.GetUser)
	}
}
