package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

// Exist checks if the given path exist (200 OK) or not (404 Not Found).
func Exist(c echo.Context) error {
	c.Set("handler_method", "Exist")

	path := filepath.Join(config.Cfg.Workspace, c.Param("*"))
	if !config.Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	if !config.Engine.Exist(path) {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}
