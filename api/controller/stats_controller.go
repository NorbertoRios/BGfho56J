package controller

import (
	"geometris-go/response"
	"geometris-go/stats"
	"geometris-go/storage"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//StatsController ...
type StatsController struct {
}

// GetServiceStats godoc
// @Summary Get service statistics
// @Description Returns service statistics
// @Tags device
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ServiceStatistics
// @Router /device/stats [get]
func (c *StatsController) GetServiceStats(ctx *gin.Context) {
	p := stats.NewProcessStat()
	p.Pid = int32(os.Getpid())
	if percentage, err := p.Process.CPUPercent(); err == nil {
		percentage = percentage * 10
		percentage = math.Round(percentage) / 10
		p.CPUPercent = percentage
	}
	devices := storage.Storage().Devices()
	udpCount := 0
	tcpCount := 0
	for _, d := range devices {
		switch d.Channel().Type() {
		case "udp":
			udpCount++
		case "tcp":
			tcpCount++
		}
	}
	resp := &response.ServiceStatistic{
		TotalDeviceCount:             len(devices),
		TotalCountByWorkers:          len(devices),
		UnregisteredConnectionsCount: 0,
		UDPConnectionsCount:          udpCount,
		TCPConnectionsCount:          tcpCount,
		ProcessInfo:                  p,
	}
	ctx.JSON(http.StatusOK, resp)
}
