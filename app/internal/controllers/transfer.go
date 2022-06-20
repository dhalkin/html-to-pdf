package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Transfer(c echo.Context) error {

	return c.String(http.StatusOK, "OK")
}
