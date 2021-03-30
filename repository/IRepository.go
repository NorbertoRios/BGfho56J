package repository

//IRepository interface for all repositories
type IRepository interface {
	Save(...interface{}) error
	Load(string) interface{}
}
