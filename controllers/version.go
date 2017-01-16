package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mdouchement/lss/config"
)

// Version returns the current version of LSS.
func Version(c echo.Context) error {
	c.Set("handler_method", "Version")

	return c.JSON(http.StatusOK, echo.Map{
		"version": config.Cfg.Version,
	})
}
