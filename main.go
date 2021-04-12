package main

import (
	"geometris-go/repository"
)

var instanse *ServiceInstanse

func main() {
	_mysqlRepo := repository.NewConsoleRepository("mysql")
	_rabbitPero := repository.NewConsoleRepository("rabbit")
	instanse = NewService(10, _mysqlRepo, _rabbitPero)
	instanse.AddServer("172.16.0.44", 10064)
	instanse.Start()
}
