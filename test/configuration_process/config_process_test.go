package test

import (
	"fmt"
	"geometris-go/core/processes/configuration"
	"geometris-go/core/processes/configuration/request"
	"geometris-go/core/processes/message"
	messageStates "geometris-go/core/processes/message/states"
	"geometris-go/core/processes/synchronization"
	synchStates "geometris-go/core/processes/synchronization/states"
	"geometris-go/core/usecase"
	"geometris-go/message/factory"
	"geometris-go/rabbitlogger"
	"geometris-go/repository"
	"geometris-go/storage"
	"geometris-go/test/mock"
	"geometris-go/unitofwork"
	"testing"
	"time"
)

var mockMysqlRepo repository.IRepository = mock.NewRepository()
var mockRabbitRepo repository.IRepository = mock.NewRepository()

func SetUP() {
	messageFactory := factory.New()
	message := messageFactory.BuildMessage([]byte("87A110550003,F001,OFF_PERIODIC,1616773466,48.746404,37.591212,33,9,0,40,0,310,0.0,4,,0,0,,,,,,,0:0,,0,0,"))
	mock.NewDeviceBuilder(message, nil, make(map[string]string), unitofwork.New(mockMysqlRepo, mockRabbitRepo)).Build()
	rabbitlogger.BuildRabbitLogger(mock.NewRepository())
}

func TestConfigWorkflow(t *testing.T) {
	SetUP()
	commands := []string{
		"1=1440;",
		"2=168.62.211.173:10100;",
		"3=10.0.0.127;",
		"4=10.0.0.127;",
		"6=device;",
		"7=device;",
		"12=65.28.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82;",
		"13=test@geometris.com;",
	}
	request := request.NewConfigRequest("geometris_87A110550003", "cfg05042021", commands)
	device := storage.Storage().Device("geometris_87A110550003")
	if device == nil {
		t.Error("Device is not nil")
	}
	process := device.Processes().Configuration()
	if process == nil {
		t.Error("Process is not nil")
	}
	mockMysqlRepo := mock.NewRepository()
	mockRabbitRepo := mock.NewRepository()
	APIUseCase := usecase.NewAPIRequestUseCase(mockMysqlRepo, mockRabbitRepo)
	APIUseCase.Launch(request, device, process)
	messageUseCase := usecase.NewUDPMessageUseCase(mockMysqlRepo, mockRabbitRepo)
	messageFactory := factory.New()
	for _, command := range commands {
		time.Sleep(500 * time.Millisecond)
		ackMessage := messageFactory.BuildMessage([]byte(fmt.Sprintf("87A110550003 ACK <SETPARAMS %v ACK>", command)))
		messageUseCase.Launch(ackMessage, nil)
	}
	processes := storage.Storage().Device("geometris_87A110550003").Processes().All()
	for _, process := range processes {
		switch process.(type) {
		case *synchronization.Process:
			{
				if process.(*synchronization.Process).Current() != nil {
					t.Error("Current synch state should be done")
				}
			}
		case *message.Process:
			{
				_, s := process.(*message.Process).Current().State().(*messageStates.InProgress)
				if !s {
					t.Error("Current message state should be done")
				}
			}
		case *configuration.Process:
			{
				if process.(*configuration.Process).Current() != nil {
					t.Error("Current config state should be done")
				}
			}
		default:
			{
				t.Error(fmt.Sprintf("Unexpected process %T", process))
			}
		}
	}
	difCRCMessage := messageFactory.BuildMessage([]byte("87A110550003,F002,OFF_PERIODIC,1616773466,48.746404,37.591212,33,9,0,40,0,310,0.0,4,,0,0,,,,,,,0:0,,0,0,"))
	messageUseCase = usecase.NewUDPMessageUseCase(mockMysqlRepo, mockRabbitRepo)
	messageUseCase.Launch(difCRCMessage, nil)
	processes = storage.Storage().Device("geometris_87A110550003").Processes().All()
	for _, process := range processes {
		switch process.(type) {
		case *synchronization.Process:
			{
				_, s := process.(*synchronization.Process).Current().State().(*synchStates.InProgress)
				if !s {
					t.Error("Current synch state should be InProgress")
				}
			}
		}
	}
	for i := 0; i < 5; i++ {
		messageUseCase.Launch(difCRCMessage, nil)
	}
	paramCRCMessage := messageFactory.BuildMessage([]byte([]byte("87A110550003 PARAMETERS 12=65.28.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82;")))
	messageUseCase = usecase.NewUDPMessageUseCase(mockMysqlRepo, mockRabbitRepo)
	messageUseCase.Launch(paramCRCMessage, nil)
	processes = storage.Storage().Device("geometris_87A110550003").Processes().All()
	for _, process := range processes {
		switch process.(type) {
		case *synchronization.Process:
			{
				if process.(*synchronization.Process).Current() != nil {
					t.Error("Current synch state should be done")
				}
			}
		case *message.Process:
			{
				_, s := process.(*message.Process).Current().State().(*messageStates.InProgress)
				if !s {
					t.Error("Current message state should be in progress.")
				}
			}
		}
	}
}
