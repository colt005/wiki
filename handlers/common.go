package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type HTTPHandler func(c echo.Context) error

func Make(h HTTPHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := h(c); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", c.Path())
			return err
		}

		return nil
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
