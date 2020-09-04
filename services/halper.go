package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// GetRouter is helper function to create a router during testing
func GetRouter(withTemplates bool) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	if withTemplates {
		router.LoadHTMLGlob("./../templates/HTML/*")
		router.Static("css", "./../templates/CSS")
		router.Static("img", "./../templates/img")
		router.Use(setUserStatus())
	}
	return router
}

// TestHTTPResponse is a helper function to process a request and test its response
func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// setUserStatus is a middleware to sets whether the user is logged in or not during testing.
func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
