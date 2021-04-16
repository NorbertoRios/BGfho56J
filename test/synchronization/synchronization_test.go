package test

import (
	"fmt"
	"geometris-go/core/processes/configuration/request"
	locationStates "geometris-go/core/processes/message/states"
	"geometris-go/core/processes/states"
	"geometris-go/core/usecase"
	"geometris-go/message/factory"
	"geometris-go/repository"
	"geometris-go/storage"
	"geometris-go/test/mock"
	"geometris-go/unitofwork"
	"os"
	"testing"
)

var mockMysqlRepo repository.IRepository = mock.NewRepository()
var mockRabbitRepo repository.IRepository = mock.NewRepository()

func SetUP(_syncParam string) {
	messageFactory := factory.New()
	message := messageFactory.BuildMessage([]byte("87A110550003,F001,OFF_PERIODIC,1616773466,48.746404,37.591212,33,9,0,40,0,310,0.0,4,,0,0,,,,,,,0:0,,0,0,."))
	mock.NewDeviceBuilder(message, nil, make(map[string]string), unitofwork.New(mockMysqlRepo, mockRabbitRepo)).Build()
}

func TestSynchronizationProcess(t *testing.T) {
	SetUP("")
	device := storage.Storage().Device("geometris_87A110550003")
	process := device.Processes().LocationRequest()
	if process != nil {
		t.Error("Process should be nil")
	}
}

func TestSynchronizationComplete(t *testing.T) {
	SetUP("")
	messageUseCase := usecase.NewUDPMessageUseCase(mockMysqlRepo, mockRabbitRepo)
	messageFactory := factory.New()
	ackMessage := messageFactory.BuildMessage([]byte("87A110550003 PARAMETERS 12=65.28.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82;"))
	messageUseCase.Launch(ackMessage, nil)
	device := storage.Storage().Device("geometris_87A110550003")
	process := device.Processes().LocationRequest()
	if process == nil {
		t.Error("Process cant be nil")
	}
}

func TestMultisync(t *testing.T) {
	os.Args[1] = ".."
	SetUP("12=65.28.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82;")
	commands := []string{
		"1=1440;",
		"2=168.62.211.173:10100;",
		"3=10.0.0.127;",
		"4=10.0.0.127;",
		"6=device;",
		"7=device;",
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
	ackMessage := messageFactory.BuildMessage([]byte(fmt.Sprintf("87A110550003 ACK <SETPARAMS %v ACK;>", commands[0])))
	messageUseCase.Launch(ackMessage, nil)
	processes := storage.Storage().Device("geometris_87A110550003").Processes().All()
	for _, process := range processes {
		if process.Symbol() == "location" {
			currentTask := process.Current()
			if currentTask != nil {
				switch currentTask.State().(type) {
				case *states.Pause:
					{

					}
				default:
					{
						t.Error(fmt.Sprintf("Unexpected task state %T", currentTask))
					}
				}
			}
		}
	}
	for i := 1; i < len(commands); i++ {
		ackMessage := messageFactory.BuildMessage([]byte(fmt.Sprintf("87A110550003 ACK <SETPARAMS %v ACK;>", commands[i])))
		messageUseCase.Launch(ackMessage, nil)
	}
	for _, process := range processes {
		if process.Symbol() == "location" {
			currentTask := process.Current()
			if currentTask != nil {
				switch currentTask.State().(type) {
				case *locationStates.InProgress:
					{
					}
				default:
					{
						t.Error(fmt.Sprintf("Unexpected task state %T", currentTask.State()))
					}
				}
			}
		}
	}
	message := messageFactory.BuildMessage([]byte("87A110550003,F001,OFF_PERIODIC,1616773466,48.746404,37.591212,33,9,0,40,0,310,0.0,4,,0,0,,,,,,,0:0,,0,0,."))
	messageUseCase.Launch(message, nil)
}
