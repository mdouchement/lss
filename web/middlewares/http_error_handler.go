package middlewares

import (
	"github.com/labstack/echo"
	"github.com/mdouchement/lss/errors"
)

// HTTPErrorHandler is a middleware that formats rendered errors.
func HTTPErrorHandler(err error, c echo.Context) {
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			c.NoContent(errors.StatusCode(err))
		} else {
			c.JSON(errors.StatusCode(err), err)
		}
	}
}
