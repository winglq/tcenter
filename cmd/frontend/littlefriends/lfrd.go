package main

import (
	"fmt"
	"os"
	"tcenter/cmd/util"
	"tcenter/common/server"
	"tcenter/common/signal"
	"tcenter/frontend/littlefriends"
	"time"

	logging "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewDefaultLittleFriendsCommand() *cobra.Command {
	cfg := &server.ServerConfig{}
	f := util.NewServerFactory(cfg)
	command := NewLittleFriendsCommand(f)
	return command
}

func NewLittleFriendsCommand(f util.Factory) *cobra.Command {
	ops := &littlefriends.LittleFriendsOptions{}
	cmds := &cobra.Command{
		Use:   "lfrd",
		Short: "Little Friends Server",
		Run: func(cmd *cobra.Command, args []string) {
			stopCh := make(chan struct{}, 1)
			signal.RegisterSignal(stopCh)
			c := littlefriends.NewService(ops)
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
	cmds.Flags().StringVar(&ops.StaticPath, "static_dir", "/tmp", "Static files dir")
	cmds.Flags().StringVar(&ops.TemplatePath, "temp_dir", "./resources/frontend/littlefriends/templates", "Templates dir path")
	return cmds
}

func Execute() {
	command := NewDefaultLittleFriendsCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
