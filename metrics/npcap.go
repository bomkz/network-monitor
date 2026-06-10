package metrics

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func ensureNpCap() {
	downloadNpCap()
	cmd := exec.Command("./npcap.exe")
	if err := cmd.Start(); err != nil {
		log.Fatal("Failed to launch NpCap installer: ", err)
	}
	log.Println("NpCap installer launched. Please complete installation and rerun netdiag.")
	os.Exit(1)
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
