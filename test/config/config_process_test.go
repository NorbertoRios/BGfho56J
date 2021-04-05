package test

import (
	"geometris-go/core/device"
	"geometris-go/core/processes/configuration/request"
	"geometris-go/core/sensors"
	"geometris-go/core/usecase"
	"geometris-go/storage"
	"geometris-go/test/mock"
	"testing"
)

func TestConfigWorkflow(t *testing.T) {
	storage := storage.Storage()
	commands := []string{
		"1=1440",
		"2=168.62.211.173:10100",
		"3=10.0.0.127",
		"4=10.0.0.127",
		"6=device",
		"7=device",
		"12=65.28.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72 .73.74.75.76.77.80.81.82",
		"13=test@geometris.com",
		"526=20.10.250.1600.350.1600.400.1600.1000.1000",
		"529=1440",
		"948=+19703769800",
		"949=updateclient",
	}
	request := request.NewConfigRequest("geometris_87A110550003", "cfg05042021", commands)
	storage.AddDevice(device.NewDevice("geometris_87A110550003", "", make(map[string]sensors.ISensor), nil))
	device := storage.Device("geometris_87A110550003")
	if device == nil {
		t.Error("Device is not nil")
	}
	process := device.Processes().Configuration()
	if process == nil {
		t.Error("Process is not nil")
	}
	usecase := usecase.NewAPIRequestUseCase(mock.NewRepository(), mock.NewRepository())
	usecase.Launch(request, device, process)
}
