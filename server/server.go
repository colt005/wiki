package server

import (
	"log/slog"
	"os"

	"github.com/colt005/wiki/handlers"
	"github.com/colt005/wiki/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	handlers handlers.HandlersI
	e        *echo.Echo
}

func New() (*Server, error) {

	s := service.New()

	h := handlers.New(&s)

	return &Server{
		handlers: h,
		e:        echo.New(),
	}, nil
}

func (s *Server) RegisterRoutes() {
	s.e.Any("/public/*", echo.WrapHandler(public()))

	s.e.GET("/", handlers.Make(s.handlers.HandleHome))
}

func (s *Server) Start() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("Server Started", "ListenAddress", listenAddr)
	s.e.Logger.Fatal(s.e.Start(listenAddr))
}
