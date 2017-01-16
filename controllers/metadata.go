package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

// Metadata returns the metadata for the given path.
func Metadata(c echo.Context) error {
	c.Set("handler_method", "Metadata")

	path := filepath.Join(config.Cfg.Workspace, c.Param("*"))
	if !config.Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	if !config.Engine.Exist(path) {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, config.Engine.Metadata(path))
}
