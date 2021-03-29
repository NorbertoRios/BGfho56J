package response

import "time"

//FacadeResponse ...
type FacadeResponse struct {
	Code            string    `json:"Code"`
	Comment         string    `json:"Comment"`
	ExecutedCommand string    `json:"ExecutedCommand"`
	CallbackID      string    `json:"CallbackId"`
	CreatedAt       time.Time `json:"CreatedAt"`
	Success         bool      `json:"Success"`
}
