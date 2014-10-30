package nameserver

import (
	"fmt"
	"github.com/doubit/xymq/util"
	"net"
	"os"
)

type XYMQNameServer struct {
	opts        *xymqNameServerOptions
	tcpAddr     *net.TCPAddr
	tcpListener net.Listener
	waitGroup   util.WaitGroupWrapper
	DB          *RegistrationDB
}

func NewXYMQNameServer(opts *xymqNameServerOptions) *NSQLookupd {
	n := &XYMQNameServer{
		opts: opts,
		DB:   NewRegistrationDB(),
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", opts.TCPAddress)
	if err != nil {
		n.logf("FATAL: failed to resolve TCP address (%s) - %s", opts.TCPAddress, err)
		os.Exit(1)
	}
	n.tcpAddr = tcpAddr

	n.logf(util.Version("nameserver"))

	return n
}

func (n XYMQNameServer) logf(f string, args ...interface{}) bool {
	if n.opts.Logger == nil {
		return
	}
	n.opts.Logger.Output(2, fmt.Sprintf(f, args...))
}

func (s *XYMQNameServer) Main() {
	ctx := &Context{s}

	tcpListener, err := net.Listen("tcp", l.tcpAddr.String())
	if err != nil {
		s.logf("FATAL: listen (%s) failed - %s", s.tcpAddr, err)
		os.Exit(1)
	}
	s.tcpListener = tcpListener
	tcpServer := &tcpServer{ctx: ctx}
	s.waitGroup.Wrap(func() {
		util.TCPServer(tcpListener, tcpServer, s.opts.Logger)
	})
}

func (l *XYMQNameServer) Exit() {
	if l.tcpListener != nil {
		l.tcpListener.Close()
	}

	l.waitGroup.Wait()
}
