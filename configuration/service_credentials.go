package configuration

//ServiceCredentials represents service credentials
type ServiceCredentials struct {
	MysqDeviceMasterConnectionString string
	WebAPIPort                       int
	GarbageDuration                  int
	UDPHost                          []string
	UDPPort                          []int
	WorkersCount                     int
	Rabbit                           *RabbitCredentials
	SystemExchange                   string
	FacadeCallbackExchange           string
	FacadeCallbackRoutingKey         string
	OrleansDebugRoutingKey           string
}
