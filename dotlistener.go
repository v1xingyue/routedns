package rdns

import (
	"crypto/tls"
	"fmt"

	"github.com/miekg/dns"
)

// DoTListener is a DNS listener/server for DNS-over-TLS.
type DoTListener struct {
	*dns.Server
}

var _ Listener = &DoTListener{}

// DoTListenerOptions contains options used by the DNS-over-TLS server.
type DoTListenerOptions struct {
	TLSConfig *tls.Config
}

// NewDoTListener returns an instance of either a UDP or TCP DNS listener.
func NewDoTListener(addr string, opt DoTListenerOptions, resolver Resolver) *DoTListener {
	return &DoTListener{
		Server: &dns.Server{
			Addr:      addr,
			Net:       "tcp-tls",
			TLSConfig: opt.TLSConfig,
			Handler:   listenHandler(resolver),
		},
	}
}

// Start the Dot server.
func (s DoTListener) Start() error {
	return s.ListenAndServe()
}

// Stop the server.
func (s DoTListener) Stop() error {
	return s.Shutdown()
}

func (s DoTListener) String() string {
	return fmt.Sprintf("DoT(%s)", s.Addr)
}