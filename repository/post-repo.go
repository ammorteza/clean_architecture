package repository

import (
	"github.com/ammorteza/clean_architecture/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (entity.Post, error)
	FetchAll() ([]entity.Post, error)
}
