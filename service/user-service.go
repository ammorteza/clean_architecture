package service

import (
	"errors"
	"github.com/ammorteza/clean_architecture/entity"
)

type UserService interface {
	IsValidUser(user *entity.User) error
	MigrateUser(user *entity.User) error
	ResetUser(user *entity.User) error
	RegisterUser(user *entity.User) error
	FetchUsers() (res []entity.User, err error)
}

func (s *service)MigrateUser(user *entity.User) error{
	if !s.repo.HasTable(user){
		return s.repo.CreateTable(user)
	}
	return nil
}

func (s *service)ResetUser(user *entity.User) error{
	if s.repo.HasTable(user) {
		return s.repo.DropTable(user)
	}

	return nil
}

func (s *service)RegisterUser(user *entity.User) error{
	if err := s.repo.Create(user); err != nil{
		return err
	}

	return nil
}

func (*service)IsValidUser(user *entity.User) error{
	if user == nil{
		return errors.New("user is empty!")
	}

	if user.Name == ""{
		return errors.New("the user's comment is empty!")
	}
	return nil
}

func (s *service)FetchUsers() (res []entity.User, err error){
	err = s.repo.Find(&entity.User{}, &res)
	return res, err
}