package repository

type DbRepository interface {
	HasTable(table interface{}) bool
	CreateTable(table interface{}) error
	DropTable(table interface{}) error
	AddForeignKey(model interface{}, field, dest, onDelete, onUpdate string) error
	Create(table interface{}) error
	Find(model interface{}, res interface{}) error
	PostRepository
	UserRepository
}

type PostRepository interface {
}

type UserRepository interface {
}
