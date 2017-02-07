package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

// Download a file according the given path.
func Download(c echo.Context) error {
	c.Set("handler_method", "Download")

	path := filepath.Join(config.Cfg.Workspace, c.Param("*"))
	if !config.Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	if !config.Engine.Exist(path) {
		return c.NoContent(http.StatusNotFound)
	}

	r, err := config.Engine.Reader(path)
	if err != nil {
		return err // err should be already well formatted
	}
	defer r.Close()

	c.Response().Header().
		Set(echo.HeaderContentLength, fmt.Sprintf("%d", config.Engine.Metadata(path)["size"]))

	return c.Stream(http.StatusOK, echo.MIMEOctetStream, r)
}
