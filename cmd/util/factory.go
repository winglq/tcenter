package util

import (
	"tcenter/common/server"

	"github.com/spf13/pflag"
)

type Factory interface {
	BindFlags(flags *pflag.FlagSet)
	ServerConfig() (*server.ServerConfig)
}

type ServerInfoFactory struct {
	Flags *pflag.FlagSet
	SvrCfg *server.ServerConfig
}

func NewServerFactory(conf *server.ServerConfig) Factory {
	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	f := &ServerInfoFactory{
		Flags: flags,
	}

	flags.StringVar(&conf.Addr, "addr", ":8080", "IP and port seperated by :")
	f.SvrCfg = conf
	return f
}

func (f *ServerInfoFactory) BindFlags(flags *pflag.FlagSet) {
	flags.AddFlagSet(f.Flags)
}

func (f *ServerInfoFactory) ServerConfig() (*server.ServerConfig) {
	return f.SvrCfg
}
