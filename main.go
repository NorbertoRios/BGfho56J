package main

import (
	"geometris-go/connection"
	"geometris-go/connection/controller"
	"geometris-go/repository"
	"geometris-go/worker"
)

func main() {
	_mysqlRepo := repository.NewConsoleRepository("mysql")
	_rabbitPero := repository.NewConsoleRepository("rabbit")
	workers := worker.NewWorkerPool(2, _mysqlRepo, _rabbitPero)
	controller := controller.NewRawDataController(workers)
	udpServer := connection.ConstructUDPServer("", 10064, controller)
	udpServer.Listen()
}
