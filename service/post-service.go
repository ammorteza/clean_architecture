package service

import (
	"errors"
	"github.com/ammorteza/clean_architecture/entity"
	"github.com/ammorteza/clean_architecture/repository"
	"math/rand"
)

var (
	repo repository.PostRepository
)

type service struct {}

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (entity.Post, error)
	FetchAll() ([]entity.Post, error)
}

func NewPostService(_repo repository.PostRepository) PostService{
	repo = _repo
	return &service{}
}

func (*service) Validate(post *entity.Post) error{
	if post == nil{
		return errors.New("post is empty!")
	}

	if post.Comment == ""{
		return errors.New("the post's comment is empty!")
	}
	return nil
}

func (*service)Create(post *entity.Post) (entity.Post, error){
	post.ID = rand.Intn(1000000)
	return repo.Save(post)
}

func (*service)FetchAll() ([]entity.Post, error){
	return repo.FetchAll()
}