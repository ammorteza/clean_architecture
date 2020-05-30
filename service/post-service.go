package service

import (
	"errors"
	"github.com/ammorteza/clean_architecture/entity"
)

type PostService interface {
	IsValidPost(post *entity.Post) error
	CreatePost(post *entity.Post) error
	FetchPosts() ([]entity.Post, error)
	MigratePost(post *entity.Post) error
	ResetPost(post *entity.Post) error
}

func (s service)MigratePost(post *entity.Post) error{
	if !s.repo.HasTable(post){
		if err := s.repo.CreateTable(post); err != nil{
			return err
		}
		if err := s.repo.AddForeignKey(post, "uId", "users(id)", "CASCADE", "CASCADE"); err != nil{
			return err
		}
	}

	return nil
}

func (service) IsValidPost(post *entity.Post) error{
	if post == nil{
		return errors.New("post is empty!")
	}

	if post.Comment == ""{
		return errors.New("the post's comment is empty!")
	}
	return nil
}

func (s service)CreatePost(post *entity.Post) error{
	//post.ID = rand.Intn(1000000)
	return s.repo.Create(post)
}

func (s service)FetchPosts() (res []entity.Post, err error){
	err = s.repo.Find(&entity.Post{}, &res)
	return res, err
}

func (s service)ResetPost(post *entity.Post) error{
	if s.repo.HasTable(post) {
		return s.repo.DropTable(post)
	}

	return nil
}