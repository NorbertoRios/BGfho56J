package repository

import "geometris-go/repository/models"

//IRepository interface for all repositories
type IRepository interface {
	Save(...interface{}) error
	Load(string) *models.DeviceActivity
}
