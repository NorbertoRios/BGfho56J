package response

import (
	"encoding/json"
	"geometris-go/logger"
	"time"
)

//NewFacadeResponse ...
func NewFacadeResponse(_callbackID, _code string, _success bool) *FacadeResponse {
	return &FacadeResponse{
		Code:       _code,
		CreatedAt:  time.Now().UTC(),
		CallbackID: _callbackID,
		Success:    _success,
	}
}

//FacadeResponse ...
type FacadeResponse struct {
	Code            string    `json:"Code"`
	Comment         string    `json:"Comment"`
	ExecutedCommand string    `json:"ExecutedCommand"`
	CallbackID      string    `json:"CallbackId"`
	CreatedAt       time.Time `json:"CreatedAt"`
	Success         bool      `json:"Success"`
}

func (resp *FacadeResponse) String() string {
	jResp, jerr := json.Marshal(resp)
	if jerr != nil {
		logger.Logger().WriteToLog(logger.Error, "[FacadeResponse | Marshal] Error while marshaling facade response. Error: ", jerr.Error())
		return ""
	}
	return string(jResp)
}
