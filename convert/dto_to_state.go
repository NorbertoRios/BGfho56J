package convert

import (
	"fmt"
	"geometris-go/configuration"
	"geometris-go/core/sensors"
	"geometris-go/core/sensors/builder"
	"geometris-go/dto"
	"geometris-go/logger"
)

//NewDTOToState ...
func NewDTOToState(_message string, _config *configuration.ReportConfiguration) *DTOToState {
	return &DTOToState{
		message:      _message,
		reportConfig: _config,
	}
}

//DTOToState ...
type DTOToState struct {
	reportConfig *configuration.ReportConfiguration
	message      string
}

//Convert ...
func (dts *DTOToState) Convert() []sensors.ISensor {
	result := []sensors.ISensor{}
	if dts.message == "" {
		return result
	}
	sb := builder.NewSensorBuilder()
	fields := dts.reportConfig.Fields()
	dtoMessage, err := dto.UnMarshalMessage(dts.message)
	if err != nil {
		logger.Logger().WriteToLog(logger.Info, "[DTOToState | Convert] Error while convert dto message to sensors.\nMessage: ", dts.message)
		return result
	}
	for _, field := range fields {
		if value, f := dtoMessage.GetValue(field.Name); !f {
			continue
		} else {
			for _, sens := range sb.Convert(field.Name, fmt.Sprintf("%v", value), field.ValueType) {
				result = append(result, sens)
			}
		}
	}
	return result
}
