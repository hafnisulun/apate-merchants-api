package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-merchants-api/models"
	"github.com/hafnisulun/apate-merchants-api/routes"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	models.ConnectDatabase()
}

func main() {
	router := gin.Default()
	// router.SetTrustedProxies([]string{"172.168.100.1"})
	v1 := router.Group("/v1")
	{
		routes.Merchants(v1)
	}
	router.Run()
}
