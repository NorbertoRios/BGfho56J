package main

import (
	"geometris-go/api/server"
	"geometris-go/configuration"
	"geometris-go/connection"
	"geometris-go/connection/controller"
	"geometris-go/connection/interfaces"
	"geometris-go/rabbitlogger"
	"geometris-go/repository"
	"geometris-go/worker"
)

func buildUDPServers(_credentials *configuration.ServiceCredentials, _controller interfaces.IController) []interfaces.IServer {
	servers := []interfaces.IServer{}
	for id, host := range _credentials.UDPHost {
		servers = append(servers, connection.ConstructUDPServer(host, _credentials.UDPPort[id], _controller))
	}
	return servers
}

//NewService ...
func NewService(_credentials *configuration.ServiceCredentials) *ServiceInstanse {
	rabbitPero := repository.NewRabbit(_credentials.Rabbit)
	mysqlRepo := repository.NewMySQL(_credentials.MysqDeviceMasterConnectionString)
	messageController := controller.NewRawDataController(worker.NewWorkerPool(_credentials.WorkersCount, mysqlRepo, rabbitPero))
	return &ServiceInstanse{
		mysqlRepo:  mysqlRepo,
		rabbitPero: rabbitPero,
		udpServers: buildUDPServers(_credentials, messageController),
		apiServer:  server.New(mysqlRepo, rabbitPero, _credentials.WebAPIPort),
	}
}

//ServiceInstanse ...
type ServiceInstanse struct {
	mysqlRepo  repository.IRepository
	rabbitPero repository.IRepository
	udpServers []interfaces.IServer
	apiServer  server.IServer
}

//Start ...
func (si *ServiceInstanse) Start() {
	rabbitlogger.BuildRabbitLogger(si.rabbitPero)
	for _, server := range si.udpServers {
		go server.Listen()
	}
	si.apiServer.Start()
}
