package rabbit

import (
	"fmt"
	"geometris-go/configuration"
	"geometris-go/logger"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channel *amqp.Channel

//Connection ...
func Connection(_config *configuration.RabbitCredentials) *amqp.Connection {
	if connection == nil {
		connectionStr := fmt.Sprintf("amqp://%v:%v@%v:%v/", _config.Username, _config.Password, _config.Host, _config.Port)
		conn, err := amqp.Dial(connectionStr)
		if err != nil {
			logger.Logger().WriteToLog(logger.Fatal, "[Rabbit | Connection] Error while create connection. Error", err)
			return nil
		}
		connection = conn
	}
	return connection
}

//Channel ...
func Channel(_config *configuration.RabbitCredentials) *amqp.Channel {
	if channel == nil {
		connection := Connection(_config)
		ch, err := connection.Channel()
		if err != nil {
			logger.Logger().WriteToLog(logger.Fatal, "[Rabbit | Channel] Error while create channel. Error", err)
			return nil
		}
		channel = ch
	}
	return channel
}
