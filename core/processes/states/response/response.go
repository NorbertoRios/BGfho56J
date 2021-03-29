package response

import (
	"encoding/json"
	"geometris-go/logger"
)

//NewFacadeResponse ...
func NewFacadeResponse(_callbackID, _code string, _success bool) *FacadeResponse {
	return &FacadeResponse{
		CallbackID: _callbackID,
		Code:       _code,
		Success:    _success,
	}
}

//FacadeResponse ...
type FacadeResponse struct {
	CallbackID string `json:"CallbackId"`
	Success    bool   `json:"Success"`
	Code       string `json:"Code"`
}

//String ...
func (resp *FacadeResponse) String() string {
	jResp, jerr := json.Marshal(resp)
	if jerr != nil {
		logger.Logger().WriteToLog(logger.Error, "[FacadeResponse | Marshal] Error while marshaling facade response. Error: ", jerr.Error())
		return ""
	}
	return string(jResp)
}
