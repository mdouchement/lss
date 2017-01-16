package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	. "github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

func Download(c echo.Context) error {
	c.Set("handler_method", "Download")

	path := filepath.Join(Cfg.Workspace, c.Param("*"))
	if !Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	if !Engine.Exist(path) {
		return c.NoContent(http.StatusNotFound)
	}

	r, err := Engine.Reader(path)
	if err != nil {
		return err // err should be already well formatted
	}
	defer r.Close()

	return c.Stream(http.StatusOK, filepath.Base(path), r)
}
