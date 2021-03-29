package observers

import (
	"geometris-go/adaptors/dto"
	"geometris-go/core/sensors"
	"geometris-go/logger"
)

//NewVINObserver ...
func NewVINObserver() *VINObserver {
	return &VINObserver{
		Symbol: "VIN",
	}
}

//VINObserver ...
type VINObserver struct {
	Symbol string
}

//Notify ...
func (o *VINObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildVINSensor(v.(string))
	}
	logger.Logger().WriteToLog(logger.Info, "[VINObserver | Notify] Cant find ", o.Symbol, " in dto message.")
	return nil
}
