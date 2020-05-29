package gorm

import (
	"github.com/ammorteza/clean_architecture/entity"
	"github.com/ammorteza/clean_architecture/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type repo struct {
	dbConn 		*gorm.DB
}

func New() repository.DbRepository {
	db, err := gorm.Open("mysql", "root:ca@1234@tcp(CA_db)/ca_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		log.Fatal(err)
	}
	return &repo{
		db,
	}
}

func (*repo)SavePost(post *entity.Post) error{
	return nil
}

func (r *repo)Find(model interface{}, res interface{}) error{
	return r.dbConn.Model(&model).Find(res).Error
}

func (r repo)Create(table interface{}) error{
	return r.dbConn.Create(table).Error
}

func (r repo)HasTable(table interface{}) bool{
	return r.dbConn.HasTable(table)
}

func (r repo)CreateTable(table interface{}) error{
	return r.dbConn.CreateTable(table).Error
}

func (r repo)ResetTable(table interface{}) error{
	return r.dbConn.DropTable(table).Error
}

func (r repo)DropTable(table interface{}) error {
	return r.dbConn.DropTable(table).Error
}
func (r repo)AddForeignKey(model interface{}, field, dest, onDelete, onUpdate string) error{
	return r.dbConn.Model(model).AddForeignKey(field, dest, onDelete, onUpdate).Error
}