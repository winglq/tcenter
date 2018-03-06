package main

import (
	"fmt"
	"os"
	"tcenter/cmd/util"
	"tcenter/common/server"
	"tcenter/common/signal"
	"tcenter/backend/registration"
	"time"

	logging "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewDefaultRegistrationCommand() *cobra.Command {
	cfg := &server.ServerConfig{}
	f := util.NewServerFactory(cfg)
	command := NewRegistrationCommand(f)
	return command
}

func NewRegistrationCommand(f util.Factory) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "regsvr",
		Short: "Registration Server",
		Run: func(cmd *cobra.Command, args []string) {
			stopCh := make(chan struct{}, 1)
			signal.RegisterSignal(stopCh)
			c := registration.NewService()
			svrConfig := f.ServerConfig()
			svrInfo, err := server.NewServerInfo(svrConfig)
			if err != nil {
				logging.Errorf("%v", err)
				return
			}
			svrInfo.Serve(c, time.Second*10, stopCh)
		},
	}
	f.BindFlags(cmds.PersistentFlags())
	return cmds
}

func Execute() {
	command := NewDefaultRegistrationCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
