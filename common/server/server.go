package server

import (
	"context"
	"net"
	"net/http"
	"time"
)

const (
	defaultKeepAlivePeriod = 3 * time.Minute
)

func (s *ServerInfo) Serve(handler http.Handler, shutdownTimeout time.Duration, stopCh <-chan struct{}) error {
	svr := &http.Server{
		Addr:           s.Listener.Addr().String(),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		<-stopCh
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		svr.Shutdown(ctx)
		cancel()
	}()
	var listener net.Listener
	listener = tcpKeepAliveListener{s.Listener.(*net.TCPListener)}
	err := svr.Serve(listener)
	return err
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
//
// Copied from Go 1.7.2 net/http/server.go

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(defaultKeepAlivePeriod)
	return tc, nil
}
