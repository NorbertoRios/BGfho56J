package test

import (
	"geometris-go/message/factory"
	"geometris-go/repository"
	"geometris-go/storage"
	"geometris-go/test/mock"
	"geometris-go/unitofwork"
	"testing"
)

var mockMysqlRepo repository.IRepository = mock.NewRepository()
var mockRabbitRepo repository.IRepository = mock.NewRepository()

func SetUP() {
	messageFactory := factory.New()
	message := messageFactory.BuildMessage([]byte("87A110550003,F001,OFF_PERIODIC,1616773466,48.746404,37.591212,33,9,0,40,0,310,0.0,4,,0,0,,,,,,,0:0,,0,0,."))
	mock.NewDeviceBuilder(message, nil, "", unitofwork.New(mockMysqlRepo, mockRabbitRepo)).Build()
}

func TestSynchronizationProcess(t *testing.T) {
	SetUP()
	device := storage.Storage().Device("geometris_87A110550003")
	process := device.Processes().LocationRequest()
	if process != nil {
		t.Error("Process should be nil")
	}
}
