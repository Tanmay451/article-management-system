package unit_test

import (
	"ams/handlers"
	"ams/models"
	"ams/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to the login page returns the login page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := services.GetRouter(true)

	r.GET("/", handlers.ShowIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	services.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a GET request to the home page returns the list of articles
// in JSON format when the Accept header is set to application/json
func TestArticleListJSON(t *testing.T) {
	r := services.GetRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/", handlers.ShowIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	services.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is JSON which can be converted to
		// an array of Article structs
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []models.Article
		err = json.Unmarshal(p, &articles)

		return err == nil && len(articles) >= 2 && statusOK
	})
}
