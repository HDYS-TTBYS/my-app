package main

import (
	"app/handlers"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

	fmt.Println()
}
