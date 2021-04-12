package test

import (
	"encoding/json"
	"fmt"
	"geometris-go/configuration"
	"geometris-go/convert"
	"geometris-go/core/device"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"geometris-go/message/factory"
	"geometris-go/message/types"
	"geometris-go/parser"
	pTypes "geometris-go/types"
	"reflect"
	"testing"
)

func TestConvertor(t *testing.T) {
	messageFactory := factory.New()
	rawMessage := messageFactory.BuildMessage([]byte{0x38, 0x37, 0x41, 0x31, 0x31, 0x30, 0x35, 0x35, 0x30, 0x30, 0x30, 0x33, 0x2C, 0x35, 0x32, 0x41, 0x33, 0x2C, 0x39, 0x33, 0x38, 0x36, 0x2C, 0x4F, 0x46, 0x46, 0x5F, 0x50, 0x45, 0x52, 0x49, 0x4F, 0x44, 0x49, 0x43, 0x2C, 0x31, 0x36, 0x31, 0x37, 0x31, 0x31, 0x31, 0x31, 0x32, 0x37, 0x2C, 0x30, 0x2E, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2C, 0x30, 0x2E, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2C, 0x33, 0x35, 0x31, 0x2C, 0x39, 0x39, 0x39, 0x2C, 0x30, 0x2C, 0x33, 0x35, 0x37, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x30, 0x2E, 0x30, 0x2C, 0x30, 0x2C, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x30, 0x3A, 0x30, 0x2C, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x00})
	messageParser := parser.NewWithDiffConfig("..", "/config/initialize/ReportConfiguration.xml")
	parsedMessage := messageParser.Parse(rawMessage, "12=28.60.65.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82")
	state := device.NewSensorBasedState([]sensors.ISensor{})
	conv := convert.NewStateToDTO(device.NewStateBasedState(state, parsedMessage.(*types.LocationMessage).Sensors()))
	dtoMessage := conv.Convert()
	strMess, jerr := json.Marshal(dtoMessage)
	if jerr != nil {
		t.Error("Error while marshaling")
	}
	should := "{\"sid\":0,\"Data\":{\"Ack\":1617111127,\"CRC\":\"52A3\",\"CoolantTemp\":0,\"CumulativeFuelEconomy\":0,\"CurrentFuelEconomy\":0,\"DTC\":\"0:0\",\"Duration\":357,\"EngineOdometer\":0,\"EngineSpeed\":0,\"FenceID\":0,\"FormatChecksum\":\"9386\",\"FuelLevel\":0,\"GpsValidity\":0,\"Heading\":0,\"IgnitionDuration\":0,\"IgnitionState\":0,\"Latitude\":0,\"LocationAge\":999,\"Longitude\":0,\"OdometerMIl\":0,\"RPM\":0,\"ReasonText\":\"OFF_PERIODIC\",\"Satellites\":0,\"SequenceID\":351,\"SpeedMPH\":0,\"ThrottlePosition\":0,\"TimeStamp\":\"2021-03-30T16:32:07Z\",\"TotalIdleDuration\":0,\"TripFuelEconomy\":0,\"VIN\":\"\"}}"
	logger.Logger().WriteToLog(logger.Info, string(strMess))
	if should != string(strMess) {
		t.Error("Unexpected DTO")
	}
}

func TestDecoder(t *testing.T) {
	dtoMessage := "{\"sid\":0,\"Data\":{\"Ack\":1617111127,\"CRC\":\"52A3\",\"CoolantTemp\":0,\"CumulativeFuelEconomy\":0,\"CurrentFuelEconomy\":0,\"DTC\":\"0:0\",\"Duration\":357,\"EngineOdometer\":0,\"EngineSpeed\":0,\"FenceID\":0,\"FormatChecksum\":\"9386\",\"FuelLevel\":0,\"GpsValidity\":0,\"Heading\":0,\"IgnitionDuration\":0,\"IgnitionState\":0,\"Latitude\":0,\"LocationAge\":999,\"Longitude\":0,\"OdometerMIl\":0,\"RPM\":0,\"ReasonText\":\"OFF_PERIODIC\",\"Satellites\":0,\"SequenceID\":351,\"SpeedMPH\":0,\"ThrottlePosition\":0,\"TimeStamp\":\"2021-03-30T16:32:07Z\",\"TotalIdleDuration\":0,\"TripFuelEconomy\":0,\"VIN\":\"\"}}"
	messageFactory := factory.New()
	rawMessage := messageFactory.BuildMessage([]byte{0x38, 0x37, 0x41, 0x31, 0x31, 0x30, 0x35, 0x35, 0x30, 0x30, 0x30, 0x33, 0x2C, 0x35, 0x32, 0x41, 0x33, 0x2C, 0x39, 0x33, 0x38, 0x36, 0x2C, 0x4F, 0x46, 0x46, 0x5F, 0x50, 0x45, 0x52, 0x49, 0x4F, 0x44, 0x49, 0x43, 0x2C, 0x31, 0x36, 0x31, 0x37, 0x31, 0x31, 0x31, 0x31, 0x32, 0x37, 0x2C, 0x30, 0x2E, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2C, 0x30, 0x2E, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2C, 0x33, 0x35, 0x31, 0x2C, 0x39, 0x39, 0x39, 0x2C, 0x30, 0x2C, 0x33, 0x35, 0x37, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x30, 0x2E, 0x30, 0x2C, 0x30, 0x2C, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x30, 0x3A, 0x30, 0x2C, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x00})
	messageParser := parser.NewWithDiffConfig("..", "/config/initialize/ReportConfiguration.xml")
	parsedMessage := messageParser.Parse(rawMessage, "12=28.60.65.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82")
	config := configuration.ReportConfig(pTypes.NewFileWithDir("..", "/config/initialize/ReportConfiguration.xml"))
	conv := convert.NewDTOToState(dtoMessage, config)
	convSensors := conv.Convert()
	sensorsMap := make(map[string]sensors.ISensor)
	for _, s := range convSensors {
		sensorsMap[s.Symbol()] = s
	}
	locationMessageSensors := parsedMessage.(*types.LocationMessage).Sensors()
	for _, s := range locationMessageSensors {
		if s.Symbol() == "Ack" {
			continue
		}
		if value, f := sensorsMap[s.Symbol()]; f {
			if !reflect.DeepEqual(value, s) {
				t.Error(fmt.Sprintf("Unexpected sensors value dto: %v; parsed: %v", value, s))
			}
		} else {
			t.Error(fmt.Sprintf("Cant find sensor by key: %v", s.Symbol()))
		}
	}
}
