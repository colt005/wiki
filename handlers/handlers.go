package handlers

import (
	"github.com/colt005/wiki/service"
)

type WikiHandler struct {
	service *service.ServiceI
}

func New(service *service.ServiceI) HandlersI {
	return &WikiHandler{
		service: service,
	}
}
