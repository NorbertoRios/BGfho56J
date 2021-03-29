package repository

import "geometris-go/core/device"

//IRepository interface for all repositories
type IRepository interface {
	Save(...device.Device) error
	Load(string) interface{}
}
