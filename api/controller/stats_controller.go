package controller

import (
	"expvar"
	"fmt"
	"geometris-go/response"
	"geometris-go/stats"
	"geometris-go/storage"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//NewStatsController ...
func NewStatsController() *StatsController {
	return &StatsController{
		serviceInfo: expvar.NewMap("Service"),
	}
}

//StatsController ...
type StatsController struct {
	serviceInfo *expvar.Map
}

// MetricsHandler ...
func (c *StatsController) MetricsHandler(ctx *gin.Context) {
	c.updateServiceInfo()
	w := ctx.Writer
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("{\n"))
	first := true
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			w.Write([]byte(",\n"))
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	w.Write([]byte("\n}\n"))
	ctx.AbortWithStatus(200)
}

func (c *StatsController) updateServiceInfo() {
	storageMetrics := storage.NewMetric()
	c.serviceInfo.Set("ManagedConnections", MetricIntValue{storageMetrics.DevicesCount()})
	c.serviceInfo.Set("TotalCountByWorkers", MetricIntValue{storageMetrics.DevicesCount()})
	c.serviceInfo.Set("UnregisteredConnectionsCount", MetricIntValue{0})
	c.serviceInfo.Set("UDPConnectionsCount", MetricIntValue{storageMetrics.ConnectionsCount("udp")})
	c.serviceInfo.Set("TCPConnectionsCount", MetricIntValue{storageMetrics.ConnectionsCount("tcp")})
}

// GetServiceStats ...
func (c *StatsController) GetServiceStats(ctx *gin.Context) {
	p := stats.NewProcessStat()
	p.Pid = int32(os.Getpid())
	if percentage, err := p.Process.CPUPercent(); err == nil {
		percentage = percentage * 10
		percentage = math.Round(percentage) / 10
		p.CPUPercent = percentage
	}
	storageMetrics := storage.NewMetric()
	resp := &response.ServiceStatistic{
		TotalDeviceCount:             storageMetrics.DevicesCount(),
		TotalCountByWorkers:          storageMetrics.DevicesCount(),
		UnregisteredConnectionsCount: 0,
		UDPConnectionsCount:          storageMetrics.ConnectionsCount("udp"),
		TCPConnectionsCount:          storageMetrics.ConnectionsCount("tcp"),
		ProcessInfo:                  p,
	}
	ctx.JSON(http.StatusOK, resp)
}
