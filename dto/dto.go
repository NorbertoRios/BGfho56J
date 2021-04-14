package dto

import (
	"encoding/json"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"sync"
)

//IMessage ...
type IMessage interface {
	SetValue(key string, value interface{})
	AppendRange(data map[string]interface{})
	GetValue(string) (interface{}, bool)
	NewSID(uint64)
	Marshal() string
}

//NewMessage returns new struct of  message
func NewMessage() *Message {
	return &Message{Data: make(map[string]interface{}), mtx: &sync.Mutex{}}
}

//Message struct for parsed messages
type Message struct {
	mtx                *sync.Mutex       `json:"-"`
	SID                uint64            `json:"sid"`
	TemperatureSensors []sensors.ISensor `json:"ts,omitempty"`
	Data               map[string]interface{}
}

//NewSID ...
func (m *Message) NewSID(_id uint64) {
	m.SID = _id
}

//GetValue from Data field
func (m *Message) GetValue(key string) (interface{}, bool) {
	m.mtx.Lock()
	value, found := m.Data[key]
	m.mtx.Unlock()
	return value, found
}

//SetValue to Data field
func (m *Message) SetValue(key string, value interface{}) {
	m.Data[key] = value
}

//Marshal append data fields to current Data
func (m *Message) Marshal() string {
	jMess, jErr := json.Marshal(m)
	if jErr != nil {
		logger.Logger().WriteToLog(logger.Error, "[Message | Marshal] Error while marshaling dto. ", jErr)
		jMess = []byte{}
	}
	return string(jMess)
}

//AppendRange append data fields to current Data
func (m *Message) AppendRange(data map[string]interface{}) {
	for k, v := range data {
		m.SetValue(k, v)
	}
}

//UnMarshalMessage given string to Message struct
func UnMarshalMessage(str string) (*Message, error) {
	message := &Message{mtx: &sync.Mutex{}}
	err := json.Unmarshal([]byte(str), message)
	if err != nil {
		return &Message{mtx: &sync.Mutex{}}, err
	}
	return message, err
}
