package main

import (
	"geometris-go/configuration"
	"geometris-go/logger"
	"geometris-go/types"
)

var instanse *ServiceInstanse

func main() {
	serviceConfig := configuration.ConstructCredentialsJSONProvider(types.NewFile("/config/initialize/credentials.example.json"))
	credentials, err := serviceConfig.ProvideCredentials()
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "[Main] Error while read credentials. Error: ", err.Error())
	}
	instanse = NewService(credentials)
	instanse.Start()
}
