package main

import (
	"geometris-go/connection"
	"geometris-go/connection/controller"
	"geometris-go/connection/interfaces"
	"geometris-go/logger"
	"geometris-go/repository"
	"geometris-go/worker"
)

//NewService ...
func NewService(_workerCount int, _mysqlRepo, _rabbitPero repository.IRepository) *ServiceInstanse {
	return &ServiceInstanse{
		udpServers:       []interfaces.IServer{},
		serverController: controller.NewRawDataController(worker.NewWorkerPool(_workerCount, _mysqlRepo, _rabbitPero)),
	}
}

//ServiceInstanse ...
type ServiceInstanse struct {
	udpServers       []interfaces.IServer
	serverController interfaces.IController
}

//AddServer ...
func (si *ServiceInstanse) AddServer(_host string, _port int) {
	si.udpServers = append(si.udpServers, connection.ConstructUDPServer(_host, _port, si.serverController))
}

//Start ...
func (si *ServiceInstanse) Start() {
	for _, server := range si.udpServers {
		server.Listen()
	}
	logger.Logger().WriteToLog(logger.Info, "[ServiceInstanse | Start] Servers are started")
}
