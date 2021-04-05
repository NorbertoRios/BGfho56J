package main

import (
	"geometris-go/connection"
	"geometris-go/connection/controller"
	"geometris-go/core/device"
	"geometris-go/core/sensors"
	"geometris-go/repository"
	"geometris-go/storage"
	"geometris-go/worker"
)

func main() {
	_mysqlRepo := repository.NewConsoleRepository("mysql")
	_rabbitPero := repository.NewConsoleRepository("rabbit")
	workers := worker.NewWorkerPool(2, _mysqlRepo, _rabbitPero)
	controller := controller.NewRawDataController(workers)
	udpServer := connection.ConstructUDPServer("", 10064, controller)
	storage := storage.Storage()
	storage.AddDevice(device.NewDevice("geometris_87A110550003", "12=28.60.65.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82", make(map[string]sensors.ISensor), nil))
	udpServer.Listen()
}
