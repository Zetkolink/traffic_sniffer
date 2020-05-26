package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	device := flag.String("device", "lo", "What device traffic to sniff.")

	flag.Parse()

	srv, err := NewSniffer(*device)

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
