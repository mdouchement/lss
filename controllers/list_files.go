package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	. "github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

func ListFiles(c echo.Context) error {
	c.Set("handler_method", "ListFiles")

	path := filepath.Join(Cfg.Workspace, c.Param("*"))
	if !Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	if !Engine.Exist(path) {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, Engine.ListFiles(path))
}
