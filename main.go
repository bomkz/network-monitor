package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iliasgal/network-monitor/pkg/db"
	"github.com/iliasgal/network-monitor/pkg/metrics"
)

func main() {
	// Setup channel to listen for termination signals
	signals := make(chan os.Signal, 1)
	// Notify for SIGINT and SIGTERM
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Start packet capture in its own goroutine
	go metrics.PacketCapture()

	host := "google.com"
	count := 4
	ticker := time.NewTicker(5 * time.Second) // Ping every 5 seconds

	for {
		select {
		case <-ticker.C:
			pingStats, err := metrics.PingHost(host, count)
			if err != nil {
				log.Fatal(err)
				return
			}
			db.WritePingMetricsToDB(pingStats)
		case info := <-metrics.PacketInfoChan:
			db.WritePacketInfoToDB(info)
		case <-signals:
			log.Println("Termination signal received, closing resources.")
			db.CloseDBClient()
			ticker.Stop()
			return
		}
	}
}
