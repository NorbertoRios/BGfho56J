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
	ginEngine.POST("/device/update_config", controller.NewDeviceController(_mysql, _rabbit).UpdateConfig)
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
