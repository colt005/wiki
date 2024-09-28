package service

import "fmt"

type WikiService struct {
}

func New() ServiceI {
	return &WikiService{}
}

func (ws *WikiService) Hello() {
	fmt.Println("Hello")
}
