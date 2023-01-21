package main

import (
	"log"
)

func main() {
	cfg, err := Configure()
	if err != nil {
		log.Fatal("Error configuring the server:", err)
	}
	log.Println("Constructing a server with the following configuration:", cfg)
	err = RunServer(cfg.BindAddr, cfg.Protocol)
	if err != nil {
		log.Fatal("Error running the server:", err)
	}
}
