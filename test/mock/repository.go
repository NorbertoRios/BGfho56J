package mock

import (
	"geometris-go/repository"
	"geometris-go/repository/models"
)

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
func (r *Repository) Load(_identity string) *models.DeviceActivity {
	return &models.DeviceActivity{}
}
