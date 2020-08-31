package handlers

import (
	"ams/models"

	"github.com/gin-gonic/gin"
)

// ShowIndexPage will handle home page
func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}
