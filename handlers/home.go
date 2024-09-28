package handlers

import (
	"github.com/colt005/wiki/views/home"
	"github.com/labstack/echo/v4"
)

func (w *WikiHandler) HandleHome(c echo.Context) error {
	return Render(c.Response().Writer, c.Request(), home.Index())
}
