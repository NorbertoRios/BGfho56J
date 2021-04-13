package main

import (
	"geometris-go/configuration"
	"geometris-go/logger"
	"geometris-go/rabbitlogger"
	"geometris-go/repository"
	"geometris-go/types"
)

var instanse *ServiceInstanse

func main() {
	serviceConfig := configuration.ConstructCredentialsJSONProvider(types.NewFile("/config/initialize/credentials.example.json"))
	credentials, err := serviceConfig.ProvideCredentials()
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "[Main] Error while read credentials. Error: ", err.Error())
	}
	rabbitPero := repository.NewRabbit(credentials.Rabbit)
	rabbitlogger.BuildRabbitLogger(rabbitPero)
	mysqRepo := repository.NewMySQL(credentials.MysqDeviceMasterConnectionString)
	instanse = NewService(credentials.WorkersCount, rabbitPero, mysqRepo)
	instanse.AddServer(credentials.UDPHost, credentials.UDPPort)
	instanse.Start()
}
