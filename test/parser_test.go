package test

import (
	"geometris-go/message/factory"
	message "geometris-go/message/types"
	"geometris-go/parser"
	"testing"
)

func TestMessageParser(t *testing.T) {
	messageFactory := factory.New()
	rawMessage := messageFactory.BuildMessage([]byte{0x38, 0x37, 0x41, 0x31, 0x31, 0x30, 0x35, 0x35, 0x30, 0x30, 0x30, 0x33, 0x2C, 0x35, 0x32, 0x41, 0x33, 0x2C, 0x39, 0x33, 0x38, 0x36, 0x2C, 0x4F, 0x46, 0x46, 0x5F, 0x50, 0x45, 0x52, 0x49, 0x4F, 0x44, 0x49, 0x43, 0x2C, 0x31, 0x36, 0x31, 0x37, 0x31, 0x31, 0x31, 0x31, 0x32, 0x37, 0x2C, 0x30, 0x2E, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2C, 0x30, 0x2E, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2C, 0x33, 0x35, 0x31, 0x2C, 0x39, 0x39, 0x39, 0x2C, 0x30, 0x2C, 0x33, 0x35, 0x37, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x30, 0x2E, 0x30, 0x2C, 0x30, 0x2C, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x2C, 0x30, 0x3A, 0x30, 0x2C, 0x2C, 0x30, 0x2C, 0x30, 0x2C, 0x00})
	messageParser := parser.NewWithDiffConfig("..", "/config/initializer/ReportConfiguration.xml")
	parsedMessage := messageParser.Parse(rawMessage, "12=28.60.65.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82")
	if parsedMessage == nil {
		t.Error("Parsed message cant be nil")
	}
	ack := parsedMessage.(*message.LocationMessage).Ack()
	if ack == "ACK 1617098255" {
		t.Error("Unexpected ack value")
	}
}
