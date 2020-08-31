package routes

import "github.com/gin-gonic/gin"

// InitializeRoutes will set up routes
func InitializeRoutes() *gin.Engine {

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())
	return router
}
