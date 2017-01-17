package controllers

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

// Upload stores the file to the given path.
func Upload(c echo.Context) error {
	c.Set("handler_method", "Upload")

	path := filepath.Join(config.Cfg.Workspace, c.Param("*"))
	if !config.Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	config.Engine.MkdirAllWithFilename(path)

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	w, err := config.Engine.Writer(path)
	if err != nil {
		return err // err should be already well formatted
	}
	defer w.Close()

	// Copy
	buf := make([]byte, 32768) // 32KB
	if _, err = io.CopyBuffer(w, src, buf); err != nil {
		return errors.NewControllersError("copy", errors.M{
			"reason": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
