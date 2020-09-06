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
	router.LoadHTMLGlob("templates/HTML/*")

	// Static serves files from the given file system root.
	// So, the relative path of file changes to css insted of http.FileServer
	router.Static("css", "templates/CSS")
	router.Static("img", "templates/img")

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", EnsureNotLoggedIn(), handlers.PerformLogin)

	}

	return router
}
