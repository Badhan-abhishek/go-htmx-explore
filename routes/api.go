package routes

import (
	"com.backend/http/api/controller"
	"com.backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeAPIRoute(r *gin.Engine) {
	r.Use(middleware.ErrorHandler())
	public := r.Group("/api")
	{
		public.POST("/sign-in", controller.Signin)
		public.POST("/sign-up", controller.Signup)
	}

	private := r.Group("/api")
	private.Use(middleware.AuthHandler())
	{
		private.GET("/get-users", controller.GetAllUsers)
		private.GET("/get-user/:email", controller.GetUser)
		private.GET("/me", controller.GetCurrentUser)
		private.PUT("/me/update", controller.UpdateUser)
	}
}
