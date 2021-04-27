package location

import (
	"geometris-go/core/processes/request"
	"geometris-go/core/usecase"
	"geometris-go/message/factory"
	"geometris-go/parser"
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
	message := messageFactory.BuildMessage([]byte("87A110550003,F001,OFF_PERIODIC,1616773466,48.746404,37.591212,33,9,0,40,0,310,0.0,4,,0,0,,,,,,,0:0,,0,0,"))
	syncParams := map[string]string{
		"F001": "12=65.28.9.36.3.4.7.8.11.12.14.17.24.150.56.51.55.70.71.72.73.74.75.76.77.80.81.82",
	}
	mock.NewDeviceBuilder(message, nil, syncParams, unitofwork.New(mockMysqlRepo, mockRabbitRepo)).Build()
}

func TestLocationRequest(t *testing.T) {
	SetUP()
	request := request.NewRequest("locationRequest", "geometris_87A110550003")
	device := storage.Storage().Device("geometris_87A110550003")
	if device == nil {
		t.Error("Device is not nil")
	}
	process := device.Processes().LocationRequest()
	if process == nil {
		t.Error("Process is not nil")
	}
	APIUseCase := usecase.NewAPIRequestUseCase(mockMysqlRepo, mockRabbitRepo)
	APIUseCase.Launch(request, device, process)
	messageUseCase := usecase.NewUDPMessageUseCase(mockMysqlRepo, mockRabbitRepo, parser.NewWithDiffConfig("..", "/config/initializer/ReportConfiguration.xml"))
	messageFactory := factory.New()
	locationMessage := messageFactory.BuildMessage([]byte("87A110550003,F001,OFF_PERIODIC,1616773466,48.746404,37.591212,33,9,0,40,0,310,0.0,4,,0,0,,,,,,,0:0,,0,0,"))
	messageUseCase.Launch(locationMessage, nil)
	process = storage.Storage().Device("geometris_87A110550003").Processes().LocationRequest()
	if process.Current() != nil {
		t.Error("All processes should be end")
	}
}
