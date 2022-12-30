package main

import (
	"flag"
)

//Config - A configuration object
type Config struct {
	BindAddr string
	Protocol string
	Encoding string
}

//Configure a Server - Generate the object which will be used
func Configure() (Config, error) {
	var (
		host     string
		port     string
		proto    string
		encoding string
	)
	flag.StringVar(&host, "host", "127.0.0.1", "The host address for which to bind")
	flag.StringVar(&port, "p", ":8900", "The port on the host for which to bind")
	flag.StringVar(&proto, "proto", "tcp", "The server transfer protocol")
	flag.StringVar(&encoding, "e", "all", "Which encoding to test")
	flag.Parse()

	cfg := Config{
		BindAddr: host + port,
		Protocol: proto,
		Encoding: encoding,
	}

	return cfg, nil
}
