package interfaces

import "net"

//IServer ...
type IServer interface {
	Listen()
	Send(*net.UDPAddr, string) (int64, error)
}
