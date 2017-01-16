package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	. "github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

func Exist(c echo.Context) error {
	c.Set("handler_method", "Exist")

	path := filepath.Join(Cfg.Workspace, c.Param("*"))
	if !Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	if !Engine.Exist(path) {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}