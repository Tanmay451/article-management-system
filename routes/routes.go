package routes

import (
	"ams/handlers"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes will set up routes
func InitializeRoutes() *gin.Engine {

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)

	return router
}
