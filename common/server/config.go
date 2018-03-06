package server

import (
	"net"
)

type ServerInfo struct {
	Listener net.Listener
}

type ServerConfig struct {
	Addr string
}

func NewServerInfo(sCfg *ServerConfig) (svrInfo *ServerInfo, err error) {
	ln, err := net.Listen("tcp", sCfg.Addr)
	if err != nil {
		return
	}
	svrInfo = &ServerInfo{
		Listener: ln,
	}
	return

}
