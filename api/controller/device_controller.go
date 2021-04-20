package controller

import (
	"fmt"
	boCommandRequest "geometris-go/core/processes/bocommand/request"
	"geometris-go/core/processes/configuration/request"
	baseRequest "geometris-go/core/processes/request"
	"geometris-go/core/usecase"
	"net"

	_ "geometris-go/docs"
	"geometris-go/logger"
	"geometris-go/repository"
	"geometris-go/response"
	"geometris-go/storage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//NewDeviceController ...
func NewDeviceController(_mysql, _rabbit repository.IRepository) *DeviceController {
	return &DeviceController{
		mysql:  _mysql,
		rabbit: _rabbit,
	}
}

//DeviceController ...
type DeviceController struct {
	mysql  repository.IRepository
	rabbit repository.IRepository
}

// SendCommand godoc
// @Summary Send command to device
// @Description Enqueue command to device
// @Tags device
// @Accept  multipart/form-data
// @Produce  json
// @Param identity formData string true "identity"
// @Param command formData string true "command"
// @Param callback_id formData string true "callback_id"
// @Success 302 {object} response.FacadeResponse
// @Failure 404 {object} response.FacadeResponse
// @Router /device/command [post]
func (c *DeviceController) SendCommand(ctx *gin.Context) {
	identity := ctx.Request.PostFormValue("identity")
	callbackID := ctx.Request.PostFormValue("callback_id")
	command := ctx.Request.PostFormValue("command")
	if !storage.Storage().DeviceExist(identity) {
		resp := &response.FacadeResponse{
			CreatedAt: time.Now().UTC(),
			Success:   false,
		}
		ctx.JSON(http.StatusNotFound, resp)
	}
	device := storage.Storage().Device(identity)
	request := boCommandRequest.NewBOCommandRequest(identity, callbackID, command)
	commandUseCase := usecase.NewAPIRequestUseCase(c.mysql, c.rabbit)
	commandUseCase.Launch(request, device, device.Processes().BOCommand(command))
	resp := &response.FacadeResponse{
		CreatedAt:       time.Now().UTC(),
		Success:         true,
		ExecutedCommand: command,
		Code:            command,
	}
	ctx.JSON(http.StatusFound, resp)
}

// SendCommandDirect godoc
// @Summary Send command to IP and port using UDP protocol
// @Description Send packet to IP:port
// @Tags device
// @Accept  multipart/form-data
// @Produce  json
// @Param ip formData string true "ip"
// @Param port formData int true "port"
// @Param command formData string true "command"
// @Success 200 {object} response.FacadeResponse
// @Failure 500 {object} response.FacadeResponse
// @Router /device/directcommand [post]
func (c *DeviceController) SendCommandDirect(ctx *gin.Context) {
	ipAddress := ctx.Request.PostFormValue("ip")
	port := ctx.Request.PostFormValue("port")
	command := ctx.Request.PostFormValue("command")

	conn, err := net.Dial("udp", fmt.Sprintf("%v:%v", ipAddress, port))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}
	defer conn.Close()
	fmt.Fprintf(conn, command)
	response := &response.FacadeResponse{
		CreatedAt:       time.Now().UTC(),
		Code:            fmt.Sprintf("Address:%v:%v; Command:%v", ipAddress, port, command),
		ExecutedCommand: command,
		Success:         true,
	}
	ctx.JSON(http.StatusOK, response)
}

// GetLocateCommand godoc
// @Summary Get locate command
// @Description Returns locate command
// @Tags device
// @Accept  multipart/form-data
// @Produce  json
// @Param identity query string true "identity"
// @Success 302 {object} response.FacadeResponse
// @Router /device/locate [get]
func (c *DeviceController) GetLocateCommand(ctx *gin.Context) {
	response := &response.FacadeResponse{
		CreatedAt:       time.Now().UTC(),
		Code:            "POLLQ VIAUDP",
		ExecutedCommand: "POLLQ VIAUDP",
		Success:         true,
	}
	ctx.JSON(http.StatusFound, response)
}

// Locate godoc
// @Summary Send locate request
// @Description Enqueue location request to device
// @Tags device
// @Accept  multipart/form-data
// @Produce  json
// @Param identity formData string true "identity"
// @Param callback_id formData string true "callback_id"
// @Success 302 {object} response.FacadeResponse
// @Failure 404 {object} response.FacadeResponse
// @Router /device/locate [post]
func (c *DeviceController) Locate(ctx *gin.Context) {
	identity := ctx.Request.PostFormValue("identity")
	callbackID := ctx.Request.PostFormValue("callback_id")
	var resp *response.FacadeResponse
	if !storage.Storage().DeviceExist(identity) {
		resp = &response.FacadeResponse{
			Code:      fmt.Sprintf("Device with 'identity'=%v online", identity),
			CreatedAt: time.Now().UTC(),
			Success:   true,
		}
		ctx.JSON(http.StatusFound, resp)
		return
	}
	device := storage.Storage().Device(identity)
	req := baseRequest.NewRequest(callbackID, identity)
	process := device.Processes().LocationRequest()
	if process == nil {
		resp = &response.FacadeResponse{
			Code:      fmt.Sprintf("Location process for device 'identity'=%v is paused", identity),
			CreatedAt: time.Now().UTC(),
			Success:   false,
		}
		ctx.JSON(http.StatusNotFound, resp)
		return
	}
	usecase.NewAPIRequestUseCase(c.mysql, c.rabbit).Launch(req, device, process)
}

// DeviceOnline godoc
// @Summary Checks device is currently connected to service
// @Description Checks device by device identity
// @Tags device
// @Accept  json
// @Produce  json
// @Param identity query string true "identity"
// @Success 302 {object} response.FacadeResponse
// @Failure 404 {object} response.FacadeResponse
// @Router /device/identity_exists [get]
func (c *DeviceController) DeviceOnline(ctx *gin.Context) {
	identity := ctx.Query("identity")
	var resp *response.FacadeResponse
	if storage.Storage().DeviceExist(identity) {
		resp = &response.FacadeResponse{
			Code:      fmt.Sprintf("Device with 'identity'=%v online", identity),
			CreatedAt: time.Now().UTC(),
			Success:   true,
		}
		ctx.JSON(http.StatusFound, resp)
		return
	}
	resp = &response.FacadeResponse{
		Code:      fmt.Sprintf("Device with 'identity'=%v offline", identity),
		CreatedAt: time.Now().UTC(),
		Success:   false,
	}
	ctx.JSON(http.StatusNotFound, resp)
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
func (c *DeviceController) UpdateConfig(ctx *gin.Context) {
	identity := ctx.Request.PostFormValue("identity")
	callbackID := ctx.Request.PostFormValue("callback_id")
	commands := ctx.PostFormArray("config")
	logger.Logger().WriteToLog(logger.Info, "[Debug] ", commands)
	if !storage.Storage().DeviceExist(identity) {
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
	ctx.JSON(http.StatusOK, resp)
	request := request.NewConfigRequest(identity, callbackID, commands)
	device := storage.Storage().Device(identity)
	usecase.NewAPIRequestUseCase(c.mysql, c.rabbit).Launch(request, device, device.Processes().Configuration())
}
