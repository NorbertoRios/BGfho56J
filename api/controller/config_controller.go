package controller

import (
	"fmt"
	"geometris-go/core/processes/configuration/request"
	"geometris-go/core/usecase"

	_ "geometris-go/docs"
	"geometris-go/logger"
	"geometris-go/repository"
	"geometris-go/response"
	"geometris-go/storage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//New ...
func NewConfig(_mysql, _rabbit repository.IRepository) *Config {
	return &Config{
		mysql:  _mysql,
		rabbit: _rabbit,
	}
}

//Config ...
type Config struct {
	mysql  repository.IRepository
	rabbit repository.IRepository
}

// UpdateConfig godoc
// @Summary Send configuration to device
// @Description Enqueue configuration to device
// @Tags device
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param identity formData string true "identity"
// @Param callback_id formData string true "callback_id"
// @Param config formData []string true "config" collectionFormat(multi)
// @Success 200 {object} response.FacadeResponse
// @Failure 404 {object} response.FacadeResponse
// @Router /device/update_config [post]
func (c *Config) UpdateConfig(ctx *gin.Context) {
	identity := ctx.Request.PostFormValue("identity")
	callbackID := ctx.Request.PostFormValue("callback_id")
	commands := ctx.PostFormArray("config")
	logger.Logger().WriteToLog(logger.Info, "[Debug] ", commands)
	device := storage.Storage().Device(identity)
	if device == nil {
		resp := &response.FacadeResponse{
			Code:      fmt.Sprintf("Cant send config  to %v . Device is offline", identity),
			CreatedAt: time.Now().UTC(),
			Success:   false,
		}
		ctx.JSON(http.StatusNotFound, resp)
		return
	}
	resp := &response.FacadeResponse{
		CreatedAt: time.Now().UTC(),
		Success:   true,
	}
	ctx.JSON(http.StatusNotFound, resp)
	request := request.NewConfigRequest(identity, callbackID, commands)
	usecase.NewAPIRequestUseCase(c.mysql, c.rabbit).Launch(request, device, device.Processes().Configuration())
}
