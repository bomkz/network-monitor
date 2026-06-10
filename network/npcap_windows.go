package network

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"syscall"

	"github.com/google/gopacket/pcap"
	"golang.org/x/sys/windows"
)

func ensureNpCap() {
	_, err := pcap.FindAllDevs()
	if err == nil {
		return
	} else if err.Error() != "couldn't load wpcap.dll" {
		log.Fatal(err)
	}

	downloadNpCap()

	npCapAbs, err := filepath.Abs("./npcap.exe")
	if err != nil {
		log.Fatal(err)
	}

	if err := launchElevated(npCapAbs); err != nil {
		log.Fatal("Failed to launch NpCap installer: ", err)
	}
	log.Println("NpCap installer launched. Please complete installation and rerun netdiag.")
	os.Exit(1)
}

func launchElevated(path string) error {
	verb, err := syscall.UTF16PtrFromString("runas")
	if err != nil {
		return err
	}
	exe, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	return windows.ShellExecute(0, verb, exe, nil, nil, windows.SW_NORMAL)
}

func downloadNpCap() {
	url := "https://npcap.com/dist/npcap-1.88.exe"
	filepath := "./npcap.exe"
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Bad HTTP Status: " + resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
