package server

import (
	"crypto/tls"
	"time"
)

// OptionFn configures options of server.
type OptionFn func(*Server)

// WithTLSConfig sets tls.Config.
func WithTLSConfig(cfg *tls.Config) OptionFn {
	return func(s *Server) {
		s.tlsConfig = cfg
	}
}

// WithReadTimeout sets readTimeout.
func WithReadTimeout(readTimeout time.Duration) OptionFn {
	return func(s *Server) {
		s.readTimeout = readTimeout
	}
}

// WithWriteTimeout sets writeTimeout.
func WithWriteTimeout(writeTimeout time.Duration) OptionFn {
	return func(s *Server) {
		s.writeTimeout = writeTimeout
	}
}
