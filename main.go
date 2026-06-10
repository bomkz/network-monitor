package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iliasgal/network-monitor/db"
	"github.com/iliasgal/network-monitor/network"
	"github.com/iliasgal/network-monitor/ui"
)

func main() {
	// Setup channel to listen for termination signals
	signals := make(chan os.Signal, 1)
	// Notify for SIGINT and SIGTERM
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	network.InitNet()

	ui.InitUi()

	go func() {
		<-signals
		log.Println("Termination signal received, closing resources.")
		db.CloseDBClient()
		os.Exit(0)
	}()

}
