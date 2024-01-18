package main

import (
	"com.quizApp/models"
	"com.quizApp/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	models.ConnectDatabase()
	routes.InitializeAPIRoute(r)
	routes.InitializeTemplateRoute(r)
	r.Run()
}
