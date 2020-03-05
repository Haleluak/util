package server

import "time"

type Options struct {
	ConnectTimeOut time.Duration
	Address string
	Name string
}

type Option func(*Options)

func Newoptions(opt ...Option)  Options{
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}

	if len(opts.Address) == 0 {
		opts.Address = DefaultAddress
	}

	if len(opts.Name) == 0 {
		opts.Address = DefaultName
	}

	if opts.ConnectTimeOut == time.Duration(0) {
		opts.ConnectTimeOut = DefaultConnectTimeOut
	}

	return opts
}

// Name server name
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Address server address
func Address(a string) Option {
	return func(o *Options) {
		o.Address = a
	}
}

// ConnectTimeOut
func ConnectTimeOut(t time.Duration) Option {
	return func(o *Options) {
		o.ConnectTimeOut = t
	}
}