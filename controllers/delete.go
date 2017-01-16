package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	. "github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

func Delete(c echo.Context) error {
	c.Set("handler_method", "Delete")

	path := filepath.Join(Cfg.Workspace, c.Param("*"))
	if !Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	if !Engine.Exist(path) {
		return c.NoContent(http.StatusNotFound)
	}

	if err := Engine.Remove(path); err != nil {
		return err // err should be already well formatted
	}

	return c.NoContent(http.StatusNoContent)
}
