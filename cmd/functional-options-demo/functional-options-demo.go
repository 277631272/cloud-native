package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	s1 := NewServer("127.0.0.1", 80)
	s2 := NewServer("127.0.0.1", 80, WithProtocol("http"))
	s3 := NewServer("127.0.0.1", 80, WithProtocol("udp"), WithTimeout(300*time.Second))
	s4 := NewServer("127.0.0.1", 80, WithProtocol("tcp"), WithMaxConns(100000))

	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
	fmt.Println("s4", s4)
}

type Server struct {
	IP       string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

func NewServer(ip string, port int, options ...Option) *Server {
	server := &Server{
		IP:       ip,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

type Option func(*Server)

func WithProtocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConns(maxConns int) Option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

func WithTLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}
