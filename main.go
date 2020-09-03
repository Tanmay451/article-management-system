// main.go

package main

import (
	"ams/routes"
	"ams/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func main() {

	// Load .env file
	err := godotenv.Load()
	services.PanicIfError(err)

	router := routes.InitializeRoutes()

	// Start serving the application
	router.Run(":" + services.GetAppPORT())
}
