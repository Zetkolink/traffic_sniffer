package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	device := os.Getenv("DEVICE")

	if device == "" {
		device = "lo"
	}

	srv, err := NewSniffer(device)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = srv.Run()

	if err != nil {
		log.Fatal(err.Error())
	}

	signals := make(chan os.Signal, 1)

	signal.Notify(signals,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-signals

	srv.Stop()
}
