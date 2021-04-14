package models

import (
	"geometris-go/repository/wrapper"
	"geometris-go/types"
	"time"

	"gorm.io/gorm"
)

func buildMessageHistory(_deviceState wrapper.IDirtyStateWrapper) *MessageHistory {
	message := _deviceState.DTOMessage()
	h := &MessageHistory{}
	h.DevID = _deviceState.Identity()
	h.EntryData = _deviceState.RawData()
	h.ParsedEntryData = []byte(message.Marshal())
	if v := _deviceState.ValueByKey("TimeStamp"); v != nil {
		h.Time = v.(*types.JSONTime).Time
	}
	h.RecievedTime = time.Now().UTC()
	if v := _deviceState.ValueByKey("ReportClass"); v != nil {
		h.ReportClass = v.(string)
	}
	if v := _deviceState.ValueByKey("ReportType"); v != nil {
		h.ReportType = v.(int32)
	}
	if v, f := message.GetValue("Reason"); f {
		h.Reason = v.(int32)
	}
	if v, f := message.GetValue("Latitude"); f {
		h.Latitude = v.(float32)
	}
	if v, f := message.GetValue("Longitude"); f {
		h.Longitude = v.(float32)
	}
	if v, f := message.GetValue("Speed"); f {
		h.Speed = v.(float32)
	}
	if v, f := message.GetValue("GpsValidity"); f {
		h.ValidFix = v.(byte)
	}
	if v, f := message.GetValue("Altitude"); f {
		h.Altitude = v.(float32)
	}
	if v, f := message.GetValue("Heading"); f {
		h.Heading = v.(float32)
	}
	if v, f := message.GetValue("IgnitionState"); f {
		h.IgnitionState = v.(byte)
	}
	if v, f := message.GetValue("Odometer"); f {
		h.Odometer = v.(int32)
	}
	if v, f := message.GetValue("Satellites"); f {
		h.Satellites = v.(int32)
	}
	if v, f := message.GetValue("Supply"); f {
		h.Supply = v.(int32)
	}
	if v, f := message.GetValue("GPIO"); f {
		h.GPIO = v.(byte)
	}
	if v, f := message.GetValue("Relay"); f {
		h.Relay = v.(byte)
	}
	return h
}

//NewMessageHistory ...
func NewMessageHistory(_deviceState wrapper.IDirtyStateWrapper) *MessageHistory {
	return buildMessageHistory(_deviceState)
}

//MessageHistory struct
type MessageHistory struct {
	ID              uint64    `gorm:"column:ID;primary_key"`
	DevID           string    `gorm:"column:DevId"`
	EntryData       []byte    `gorm:"column:EntryData"`
	ParsedEntryData []byte    `gorm:"column:ParsedEntryData"`
	Time            time.Time `gorm:"column:Time"`
	RecievedTime    time.Time `gorm:"column:RecievedTime"`
	ReportClass     string    `gorm:"column:ReportClass"`
	ReportType      int32     `gorm:"column:ReportType"`
	Reason          int32     `gorm:"column:Reason"`
	Latitude        float32   `gorm:"column:Latitude"`
	Longitude       float32   `gorm:"column:Longitude"`
	Speed           float32   `gorm:"column:Speed"`
	ValidFix        byte      `gorm:"column:ValidFix"`
	Altitude        float32   `gorm:"column:Altitude"`
	Heading         float32   `gorm:"column:Heading"`
	IgnitionState   byte      `gorm:"column:IgnitionState"`
	Odometer        int32     `gorm:"column:Odometer"`
	Satellites      int32     `gorm:"column:Satellites"`
	Supply          int32     `gorm:"column:Supply"`
	GPIO            byte      `gorm:"column:GPIO"`
	Relay           byte      `gorm:"column:Relay"`
}

//MessageHistoryTable ...
func MessageHistoryTable(h *MessageHistory) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tableName := "raw_data." + h.DevID[len(h.DevID)-2:]
		return tx.Table(tableName)
	}
}
