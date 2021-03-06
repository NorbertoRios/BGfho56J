package configuration

import (
	"fmt"
	"geometris-go/logger"
	"geometris-go/types"
)

//NewReportConfig ...
func NewReportConfig(_file types.IFile) *ReportConfiguration {
	file := _file
	provider := ConstructXMLProvider(file)
	return constructReportConfiguration(provider)
}

//ReportConfiguration represents report config
type ReportConfiguration struct {
	fields []Field
}

//ConstructReportConfiguration create report config instance
func constructReportConfiguration(provider IReportConfigProvider) *ReportConfiguration {
	fields, err := provider.Provide()
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "[ReportConfiguration | constructReportConfiguration] Error while constructing report configuration. Error: ", err)
	}
	configuration := &ReportConfiguration{
		fields: fields,
	}
	return configuration
}

//Fields ...
func (reportConfiguration *ReportConfiguration) Fields() []Field {
	return reportConfiguration.fields
}

//GetFieldByID returns description for field by id
func (reportConfiguration *ReportConfiguration) GetFieldByID(id string) (*Field, error) {
	for _, reportField := range reportConfiguration.fields {
		if reportField.ID == id {
			return &reportField, nil
		}
	}
	return nil, fmt.Errorf("Not found field with id:%v", id)
}

//GetFieldsByIds returns fields array by ids
func (reportConfiguration *ReportConfiguration) GetFieldsByIds(ids []string) []*Field {
	result := make([]*Field, 0)
	for _, id := range ids {
		if id == "28" || id == "65" {
			continue
		}
		if reportField, err := reportConfiguration.GetFieldByID(id); err == nil {
			result = append(result, reportField)
		} else {
			logger.Logger().WriteToLog(logger.Error, "[GetReportColumnsByIds] ", err)
		}
	}
	return result
}
