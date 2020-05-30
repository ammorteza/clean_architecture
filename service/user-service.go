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
	UpdateUser(user *entity.User) error
	FirstUser(res *entity.User) error
	FetchUsers() (res []entity.User, err error)
}

func (s service)MigrateUser(user *entity.User) error{
	if !s.repo.HasTable(user){
		return s.repo.CreateTable(user)
	}
	return nil
}

func (s service)ResetUser(user *entity.User) error{
	if s.repo.HasTable(user) {
		return s.repo.DropTable(user)
	}

	return nil
}

func (s service)RegisterUser(user *entity.User) error{
	return s.repo.Create(user)
}

func (s service)UpdateUser(user *entity.User) error{
	return s.repo.Save(user)
}

func (service)IsValidUser(user *entity.User) error{
	if user == nil{
		return errors.New("user is empty!")
	}

	if user.Name == ""{
		return errors.New("the user's comment is empty!")
	}
	return nil
}

func (s service)FetchUsers() (res []entity.User, err error){
	err = s.repo.Find(&entity.User{}, &res)
	return
}

func (s service)FirstUser(res *entity.User) error{
	return s.repo.First(res)
}