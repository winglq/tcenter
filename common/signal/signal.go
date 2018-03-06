package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func RegisterSignal(stopCh chan<- struct{}) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		close(stopCh)
	}()
}
