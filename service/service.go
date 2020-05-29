package service

import (
	"github.com/ammorteza/clean_architecture/repository"
)

type service struct {
	repo repository.DbRepository
}

type AppService interface {
	PostService
	UserService
}

func New(_repo repository.DbRepository) AppService{
	service := &service{
		repo : _repo,
	}
	return service
}