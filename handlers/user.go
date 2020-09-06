package handlers

import (
	"github.com/gin-gonic/gin"
)

// PerformLogin is a handler for login process
func PerformLogin(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	// var sameSiteCookie http.SameSite;

	// Check if the username/password combination is valid
	if isUserValid(username, password) {
		// If the username/password is valid set the token in a cookie
		token := generateSessionToken()
		c.SetCookie("token", token, int(3600), "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		render(c, gin.H{
			"title": "Failuer"}, "index.html")
		// c.HTML(http.StatusBadRequest, "login.html", gin.H{
		// 	"ErrorTitle":   "Login Failed",
		// 	"ErrorMessage": "Invalid credentials provided"})
	}
}
