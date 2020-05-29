package controller

import (
	service2 "github.com/ammorteza/clean_architecture/service"
)

type controller struct {
	service service2.AppService
}

type AppController interface {
	postController
	userController
}

func New(service service2.AppService) AppController {
	return &controller{
		service,
	}
}