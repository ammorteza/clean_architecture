package repository

type DbRepository interface {
	HasTable(table interface{}) bool
	CreateTable(table interface{}) error
	DropTable(table interface{}) error
	AddForeignKey(model interface{}, field, dest, onDelete, onUpdate string) error
	Create(table interface{}) error
	Find(model interface{}, res interface{}) error
	First(res interface{}) error
	Save(model interface{}) error

	WithTx(tx Tx) DbRepository
	Begin() (Tx, error)
	Rollback() error
	Commit() error
	PostRepository
	UserRepository
}

type Tx struct {
	ID 				string
}

type PostRepository interface {
}

type UserRepository interface {
}
