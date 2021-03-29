package test

import (
	"geometris-go/message/factory"
	"geometris-go/message/types"
	"reflect"
	"testing"
)

func TestBuildAckMessage(t *testing.T) {
	messageFactory := factory.New()
	strMessage := "87A110550003 ACK <SETPARAMS 104=1;103=1;9=internet;10=;11= ACK;>"
	message := messageFactory.BuildMessage([]byte(strMessage))
	ackMessage, f := message.(*types.AckMessage)
	if !f {
		t.Error("Unexpected type of message")
	}
	if message.Identity() != "geometris_87A110550003" {
		t.Error("Unexpected identity")
	}
	if ackMessage.Command != "SETPARAMS 104=1;103=1;9=internet;10=;11= ACK;" {
		t.Error("Unexpected command")
	}
}

func TestBuildParameterMessage(t *testing.T) {
	messageFactory := factory.New()
	strMessage := "87A110550003 PARAMETERS 104=1.000000;12=65.28.9.36.3.4.7.8.11.12.14.17.24.150.56.51.55.70.71.72.73.74.75.76.77.80.81.82;"
	message := messageFactory.BuildMessage([]byte(strMessage))
	paramMessage, f := message.(*types.ParametersMessage)
	if !f {
		t.Error("Unexpected type of message")
	}
	if message.Identity() != "geometris_87A110550003" {
		t.Error("Unexpected identity")
	}
	if len(paramMessage.Parameters) == 0 {
		t.Error("Unexpected parameters count")
	}
	if paramMessage.Parameters["104"] != "1.000000" {
		t.Error("Unexpected parameters value")
	}
	if paramMessage.Parameters["12"] != "65.28.9.36.3.4.7.8.11.12.14.17.24.150.56.51.55.70.71.72.73.74.75.76.77.80.81.82" {
		t.Error("Unexpected parameters value")
	}
}

func TestBuildRawLocationMessage(t *testing.T) {
	messageFactory := factory.New()
	strMessage := "87A110550003,F001,OFF_PERIODIC,1616773466,48.746404,37.591212,33,9,0,40,0,310,0.0,4,,0,0,,,,,,,0:0,,0,0,"
	message := messageFactory.BuildMessage([]byte(strMessage))
	locationMessage, f := message.(*types.RawLocationMessage)
	if !f {
		t.Error("Unexpected type of message")
	}
	if message.Identity() != "geometris_87A110550003" {
		t.Error("Unexpected identity")
	}
	testSlise := []string{"F001", "OFF_PERIODIC", "1616773466", "48.746404", "37.591212", "33", "9", "0", "40", "0", "310", "0.0", "4", "", "0", "0", "", "", "", "", "", "", "0:0", "", "0", "0"}
	if !reflect.DeepEqual(testSlise, locationMessage.RawData) {
		t.Error("Unexpected raw data in message")
	}
}
