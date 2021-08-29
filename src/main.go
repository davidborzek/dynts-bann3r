package main

import (
	"bytes"
	"dynts-bann3r/src/config"
	"dynts-bann3r/src/image"
	"dynts-bann3r/src/label"
	"dynts-bann3r/src/teamspeak"
	i "image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/multiplay/go-ts3"
)

func main() {
	cfg := config.LoadConfig()

	port := os.Getenv("DYNTS_BANN3R_PORT")
	if port == "" {
		port = "9000"
	}

	client := teamspeak.Login(cfg.Connection)
	defer client.Close()

	log.Printf("[starting] dynts-bann3r - refreshing every %d seconds \n", cfg.RefreshInterval)

	go serveBanner(port, cfg.Labels)
	schedule(cfg, client)
}

var banner i.Image

var clientIpNicknameMap map[string]string

func schedule(cfg config.Config, client *ts3.Client) {
	filledLabels := make([]config.Label, len(cfg.Labels))
	copy(filledLabels, cfg.Labels)

	for {
		log.Printf("[schedule] refreshing banner.png \n")

		clientIpNicknameMap = teamspeak.RefreshClientIpNicknameMap(client)

		for i, val := range cfg.Labels {
			if !strings.Contains(val.Text, "%nickname%") {
				filledLabels[i].Text = label.GenerateLabel(val.Text, client)
			} else {
				filledLabels[i].Text = ""
			}
		}

		banner = image.AddLabelsToImage(filledLabels, cfg.TemplatePath)
		time.Sleep(time.Duration(cfg.RefreshInterval) * time.Second)
	}
}

func serveBanner(port string, labels []config.Label) {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		remoteIp := strings.Split(r.RemoteAddr, ":")[0]
		log.Println(remoteIp)

		servedBanner := banner

		nickname := clientIpNicknameMap[remoteIp]

		log.Println(nickname)

		for i, label := range labels {
			if nickname != "" && strings.Contains(label.Text, "%nickname%") {
				labels[i].Text = strings.ReplaceAll(labels[i].Text, "%nickname%", nickname)
				servedBanner = image.AddLabelToImage(labels[i], servedBanner)
			}
		}

		buffer := new(bytes.Buffer)
		err := png.Encode(buffer, servedBanner)

		if err == nil {
			rw.Write(buffer.Bytes())
		}
	})

	log.Println("[http] Serving the banner on localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
