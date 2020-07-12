package main

import (
	"atto/internal/app/daemon"
	"atto/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handleInterrupt()

	log := logger.New(logger.Info)

	d, err := daemon.New(log)
	if err != nil {
		log.Fatal("failed to create daemon", err)
		exit()
	}

	log.Info("Atto Daemon started!")

	log.Fatal("daemon crashed", d.Start())
	exit()
}

func handleInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		exit()
	}()
}

func exit() {
	os.Exit(0)
}
