package server

import (
	"fmt"
	"geometris-go/api/controller"
	"geometris-go/logger"
	"geometris-go/repository"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//New ...
func New(_mysql, _rabbit repository.IRepository, _port int) IServer {
	ginEngine := gin.Default()
	deviceController := controller.NewDeviceController(_mysql, _rabbit)
	statController := controller.NewStatsController()
	ginEngine.POST("/device/update_config", deviceController.UpdateConfig)
	ginEngine.POST("/device/command", deviceController.SendCommand)
	ginEngine.POST("/device/directcommand", deviceController.SendCommandDirect)
	ginEngine.POST("/device/locate", deviceController.Locate)
	ginEngine.GET("/device/locate", deviceController.GetLocateCommand)
	ginEngine.GET("/device/identity_exists", deviceController.DeviceOnline)

	ginEngine.GET("/device/stats", statController.GetServiceStats)
	ginEngine.GET("/debug/vars", statController.MetricsHandler)

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Server{
		server: ginEngine,
		port:   _port,
	}
}

//Server ...
type Server struct {
	server *gin.Engine
	port   int
}

//Start ...
func (s *Server) Start() {
	logger.Logger().WriteToLog(logger.Info, "[Server | Start] Http server is started on port ", s.port)
	s.server.Run(fmt.Sprintf(":%v", s.port))
}
