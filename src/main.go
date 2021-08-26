package main

import (
	"dynts-bann3r/src/config"
	"dynts-bann3r/src/image"
	"dynts-bann3r/src/label"
	"dynts-bann3r/src/teamspeak"
	"fmt"
	"log"
	"time"

	"github.com/multiplay/go-ts3"
)

func main() {
	cfg := config.LoadConfig()

	client := teamspeak.Login(cfg.Connection)
	defer client.Close()

	fmt.Printf("Starting dynts-bann3r. Updating every 5 seconds...")

	schedule(cfg, client)

}

func schedule(cfg config.Config, client *ts3.Client) {
	filledLabels := make([]config.Label, len(cfg.Labels))
	copy(filledLabels, cfg.Labels)

	for {
		fmt.Printf("updating data...\n")

		for i, val := range cfg.Labels {
			text, err := label.GenerateLabel(val.Text, client)

			if err != nil {
				log.Fatal(err)
			}

			filledLabels[i].Text = text
		}

		image.AddLabelsToImage(filledLabels, "template.png", "banner.png")

		time.Sleep(5 * time.Second)
	}
}
