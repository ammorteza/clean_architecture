package repository

import (
	"github.com/ammorteza/clean_architecture/entity"
	"math/rand"
)

type repo struct {}

func NewMysqlRepository() PostRepository {
	return &repo{}
}

func (* repo)Save(post *entity.Post) (entity.Post, error){
	return *post, nil
}

func (*repo)FetchAll() ([]entity.Post, error){
	posts := make([]entity.Post, 0)
	posts = append(posts, entity.Post{ID: rand.Intn(1000000), Comment: "test2"})
	posts = append(posts, entity.Post{ID: rand.Intn(1000000), Comment: "test3"})
	posts = append(posts, entity.Post{ID: rand.Intn(1000000), Comment: "test4"})
	return posts, nil
}