package controllers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/errors"
)

// ListFiles the directories and files of the given path.
func ListFiles(c echo.Context) error {
	c.Set("handler_method", "ListFiles")

	path := filepath.Join(config.Cfg.Workspace, c.Param("*"))
	if !config.Engine.IsPathValid(path) {
		return errors.NewControllersError("invalid_path", errors.M{
			"path": path,
		})
	}

	if !config.Engine.Exist(path) {
		return c.NoContent(http.StatusNotFound)
	}

	var depth int
	if d, ok := c.QueryParams()["depth"]; ok {
		v, err := strconv.Atoi(d[0])
		if err != nil {
			return errors.NewControllersError("invalid_query", errors.M{
				"query": "depth",
			})
		}
		depth = v
	}

	return c.JSON(http.StatusOK, config.Engine.ListFiles(path, depth))
}
