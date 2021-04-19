package connection

import (
	"fmt"
	"geometris-go/connection/interfaces"
	"geometris-go/logger"
	"net"
	"time"
)

//ConstructUDPChannel returns new channel
func ConstructUDPChannel(addr *net.UDPAddr, server *UDPServer) interfaces.IChannel {
	return &UDPChannel{
		ServerInstance: server,
		ConnectedAt:    time.Now().UTC(),
		clientAddr:     addr,
	}
}

//UDPChannel cahnnel for device
type UDPChannel struct {
	ServerInstance *UDPServer
	ConnectedAt    time.Time
	received       int64
	transmitted    int64
	clientAddr     *net.UDPAddr
}

//Received received bytes
func (c *UDPChannel) Received() int64 {
	return c.received
}

//Send message to device by UDP
func (c *UDPChannel) Send(message string) error {
	var err error
	var trs int64
	trs, err = c.ServerInstance.Send(c.clientAddr, message)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[UDPChannel | Send] Error: %v", err))
		return err
	}
	logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[UDPChannel | Send] Message(string): %v . TO: %v ", string(message), c.clientAddr.String()))
	c.AddTransmitted(trs)
	return nil
}

//Type ...
func (c *UDPChannel) Type() string {
	return "udp"
}

//Transmitted transmitted bytes
func (c *UDPChannel) Transmitted() int64 {
	return c.transmitted
}

//AddTransmitted to cahnnel
func (c *UDPChannel) AddTransmitted(count int64) {
	c.transmitted += count
}

//AddReceived to cahnnel
func (c *UDPChannel) AddReceived(count int64) {
	c.received += count
}

//RemoteAddr client address
func (c *UDPChannel) RemoteAddr() *net.UDPAddr {
	return c.clientAddr
}

//RemoteIP indicates device remote address
func (c *UDPChannel) RemoteIP() string {
	return c.clientAddr.String()
}
