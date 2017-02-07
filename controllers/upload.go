package controllers

import (
	"fmt"
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

	w, err := config.Engine.Writer(path)
	if err != nil {
		return err // err should be already well formatted
	}
	defer w.Close()

	// Stream to destination
	if err = streamParts(w, c.Request(), path); err != nil {
		return errors.NewControllersError("copy", errors.M{
			"reason": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}

func streamParts(w io.Writer, req *http.Request, path string) error {
	mr, err := req.MultipartReader()
	if err != nil {
		return err
	}

	found := false
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			if !found {
				return fmt.Errorf("http multipart: no such file")
			}
			return nil
		}
		if err != nil {
			badUploadHandler(req, path)
			return err
		}

		if p.FormName() == "file" {
			// Copy/Download
			if _, err = io.Copy(w, p); err != nil {
				badUploadHandler(req, path)
				return err
			}
			found = true
		}
	}
}

func badUploadHandler(req *http.Request, path string) {
	config.Engine.Remove(path)
}
