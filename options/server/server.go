package server

import (
	"fmt"
	"sync"
	"time"
)

var (
	DefaultAddress  =  ":0"
	DefaultName  =  "server"
	DefaultConnectTimeOut  = time.Second *  4
)

type  Server  struct {
	sync.RWMutex
	opts Options
}

func NewServer(opts ...Option)  Server{
	options := Newoptions(opts...)
	return Server{
		opts: options,
	}
}

func (s *Server) Options() Options {
	s.RLock()
	opts  := s.opts
	s.RUnlock()
	return opts
}

func (s *Server) Init(opts ...Option) error {
	s.Lock()
	for  _, opt  :=  range opts {
		opt(&s.opts)
	}
	s.Unlock()
	return  nil
}

func (s *Server) Start() error {
	fmt.Println(s.opts.Name)
	fmt.Println(s.opts.Address)
	return  nil
}

func (s *Server) Stop() error {
	return  nil
}