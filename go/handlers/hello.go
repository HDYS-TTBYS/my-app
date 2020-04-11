package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//Hello GET /api/test
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
