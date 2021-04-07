package connection

import (
	"encoding/hex"
	"fmt"
	"geometris-go/connection/controller"
	"geometris-go/logger"
	"net"
)

//ConstructUDPServer returns new UDP server
func ConstructUDPServer(host string, port int, _controller *controller.RawDataController) *UDPServer {
	addr := fmt.Sprintf("%v:%v", host, port)
	logger.Logger().WriteToLog(logger.Info, "[UDPServer] Start at: ", addr)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "[UDPServer] Wrong UDP Address: ", addr)
		return nil
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "Create udp server error: ", err.Error())
		return nil
	}

	server := &UDPServer{
		port:       port,
		connection: udpConn,
		controller: _controller,
	}
	return server
}

//UDPServer for genx service
type UDPServer struct {
	port       int
	connection *net.UDPConn
	controller *controller.RawDataController
}

//Listen incoming packet
func (server *UDPServer) Listen() {
	for {
		var buf [4096]byte
		n, addr, err := server.connection.ReadFromUDP(buf[0:])
		if err != nil {
			logger.Logger().WriteToLog(logger.Fatal, "Error Reading from udp connection: ", err.Error())
			return
		}
		logger.Logger().WriteToLog(logger.Info, "Received UDP packet:", hex.EncodeToString(buf[0:n]), ". IP: ", addr.String())
		channel := ConstructUDPChannel(addr, server)
		server.controller.Process(buf[0:n], channel)
	}
}

//SendString send string
func (server *UDPServer) SendString(addr *net.UDPAddr, packet string) (int64, error) {
	n, err := server.connection.WriteToUDP([]byte(packet), addr)
	logger.Logger().WriteToLog(logger.Info, "[UDPServer | SendString] Bytes sent. ", []byte(packet))
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[UDPServer | SendString] Error while sending string. ", err)
		return 0, err
	}
	return int64(n), nil
}

//SendBytes send bytes
func (server *UDPServer) SendBytes(addr *net.UDPAddr, packet []byte) (int64, error) {
	n, err := server.connection.WriteToUDP(packet, addr)
	logger.Logger().WriteToLog(logger.Info, "[UDPServer | SendBytes] Bytes sent. ", packet)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[UDPServer | SendBytes] Error while sending bytes. ", err)
		return 0, err
	}
	return int64(n), nil
}
