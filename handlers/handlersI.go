package handlers

import "github.com/labstack/echo/v4"

type HandlersI interface {
	HandleHome(c echo.Context) error
}
