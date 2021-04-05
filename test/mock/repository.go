package mock

import "geometris-go/repository"

//NewRepository ...
func NewRepository() repository.IRepository {
	return &Repository{}
}

//Repository ...
type Repository struct {
}

//Save ...
func (r *Repository) Save(values ...interface{}) error {
	return nil
}

//Load ...
func (r *Repository) Load(_identity string) interface{} {
	return _identity
}
