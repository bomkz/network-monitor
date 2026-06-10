package network

import (
	"log"

	"github.com/google/gopacket/pcap"
)

func InitNet() {
	ensureNpCap()
	getNetDevices()
}

func getNetDevices() {
	var err error
	NetDevs, err = pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

}
