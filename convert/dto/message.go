package dto

//Message ...
type Message struct {
	SyncParameter []string                      `json:"SyncParameter"`
	Data          map[string]interface{}        `json:"Data"`
	Sensors       map[string]*TemperatureSensor `json:"ts"`
	SID           uint64                        `json:"sid"`
}
