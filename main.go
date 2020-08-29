// main.go

package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Start serving the application
	router.Run(":" + os.Getenv("APP_PORT"))
}
