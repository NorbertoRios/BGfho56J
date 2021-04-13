package configuration

//RabbitCredentials rabbit config
type RabbitCredentials struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	DeviceStatesExchange string
	FacadeExchange       string
	FacadeRoutingKey     string
	LoggerRoutingKey     string
	Retry                int
}
