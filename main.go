package main

import (
	"com.backend/models"
	"com.backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	routes.InitializeAPIRoute(r)
	r.Run()
}
