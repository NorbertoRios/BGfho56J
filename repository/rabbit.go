package repository

import (
	"fmt"
	"geometris-go/configuration"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
	"geometris-go/rabbit"
	"geometris-go/repository/models"
	"geometris-go/repository/wrapper"

	"github.com/streadway/amqp"
)

//NewRabbit ...
func NewRabbit(_config *configuration.RabbitCredentials) IRepository {
	_connection := connection(_config)
	return &Rabbit{
		connection: _connection,
		channel:    channel(_connection),
		retry:      _config.Retry,
		config:     _config,
	}
}

//Connection ...
func connection(_config *configuration.RabbitCredentials) *amqp.Connection {
	connectionStr := fmt.Sprintf("amqp://%v:%v@%v:%v/", _config.Username, _config.Password, _config.Host, _config.Port)
	conn, err := amqp.Dial(connectionStr)
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "[Rabbit | Connection] Error while create connection. Error", err)
		return nil
	}
	return conn
}

//Channel ...
func channel(connection *amqp.Connection) *amqp.Channel {
	ch, err := connection.Channel()
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "[Rabbit | Channel] Error while create channel. Error", err)
		return nil
	}
	return ch
}

//Rabbit ...
type Rabbit struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	retry      int
	config     *configuration.RabbitCredentials
}

//Load ...
func (r *Rabbit) Load(string) *models.DeviceActivity {
	logger.Logger().WriteToLog(logger.Error, "[Rabbit | Load] Unexpected method call")
	return nil
}

//Save ...
func (r *Rabbit) Save(values ...interface{}) error {
	for _, value := range values {
		switch value.(type) {
		case []wrapper.IDirtyStateWrapper:
			{
				return r.saveDeviceState(value.([]wrapper.IDirtyStateWrapper))
			}
		case []interfaces.ITask:
			{
				return r.saveTasks(value.([]interfaces.ITask))
			}
		case string:
			{
				return r.publish(value.(string), r.config.DeviceStatesExchange, r.config.LoggerRoutingKey, r.retry)
			}
		default:
			{
				logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[Rabbit | Save] Unexpected message type %T", value))
			}
		}
	}
	return nil
}

func (r *Rabbit) saveTasks(_tasks []interfaces.ITask) error {
	for _, task := range _tasks {
		err := r.publish(task.FacadeResponse(), r.config.FacadeExchange, r.config.FacadeRoutingKey, r.retry)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Rabbit) saveDeviceState(_states []wrapper.IDirtyStateWrapper) error {
	for _, state := range _states {
		dtoMessage := state.DTOMessage()
		err := r.publish(dtoMessage.Marshal(), r.config.DeviceStatesExchange, state.Identity()[len(state.Identity())-2:], r.retry)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Rabbit) publish(message, exchange, routingKey string, retry int) error {
	if retry == 0 {
		logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[Rabbit | Publish] Message %v didnt publish. Retry is 0", message))
		return nil
	}
	err := rabbit.Channel(r.config).Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			DeliveryMode:    2,
			Headers:         make(amqp.Table, 0),
			Body:            []byte(message),
		})
	logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[Rabbit | Publish] Publish message: %v to %v:%v", message, exchange, routingKey))
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[Rabbit | Publish] Error while publish message: %v to %v:%v. Try to reconnect to rabbit", message, exchange, routingKey))
		r.reconnect()
		retry = retry - 1
		r.publish(message, exchange, routingKey, retry)
	}
	return err
}

func (r *Rabbit) reconnect() {
	r.connection.Close()
	r.channel.Close()
	r.connection = connection(r.config)
	r.channel = channel(r.connection)
}
