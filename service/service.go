package service

import (
	"github.com/ammorteza/clean_architecture/repository"
)

type service struct {
	repo repository.DbRepository
}

type AppService interface {
	BeginTx() (repository.Tx, error)
	RollbackTx() error
	CommitTx() error
	WithTx(tx repository.Tx) AppService
	PostService
	UserService
}

func New(_repo repository.DbRepository) AppService{
	service := &service{
		repo : _repo,
	}
	return service
}

func (s service)WithTx(tx repository.Tx) AppService{
	temp := s
	temp.repo = temp.repo.WithTx(tx)
	return temp
}

func (s service)BeginTx() (repository.Tx, error){
	return s.repo.Begin()
}

func (s service)RollbackTx() error{
	return s.repo.Rollback()
}

func (s service)CommitTx() error{
	return s.repo.Commit()
}