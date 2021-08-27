package main

import (
	"dynts-bann3r/src/config"
	"dynts-bann3r/src/image"
	"dynts-bann3r/src/label"
	"dynts-bann3r/src/teamspeak"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/multiplay/go-ts3"
)

func main() {
	cfg := config.LoadConfig()

	client := teamspeak.Login(cfg.Connection)
	defer client.Close()

	log.Printf("[starting] dynts-bann3r - refreshing every %d seconds \n", cfg.RefreshInterval)

	schedule(cfg, client)

}

func schedule(cfg config.Config, client *ts3.Client) {
	filledLabels := make([]config.Label, len(cfg.Labels))
	copy(filledLabels, cfg.Labels)

	go serveBanner()

	for {
		log.Printf("[schedule] refreshing banner.png \n")

		for i, val := range cfg.Labels {
			text, err := label.GenerateLabel(val.Text, client)

			if err != nil {
				log.Fatal(err)
			}

			filledLabels[i].Text = text
		}

		image.AddLabelsToImage(filledLabels, "template.png", "banner.png")

		time.Sleep(time.Duration(cfg.RefreshInterval) * time.Second)
	}
}

func serveBanner() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		dat, err := ioutil.ReadFile("banner.png")

		if err == nil {
			rw.Write(dat)
		} else {
			log.Printf("An error occurred serving the banner.png: %v \n", err)
		}
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
