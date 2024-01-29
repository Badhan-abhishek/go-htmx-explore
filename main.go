package main

import (
	"com.backend/models"
	"com.backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()

	config.AllowAllOrigins = true

	r.Use(cors.New(config))

	models.ConnectDatabase()
	routes.InitializeAPIRoute(r)
	r.Run()
}
