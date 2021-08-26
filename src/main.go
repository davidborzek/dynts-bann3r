package main

import (
	"dynts-bann3r/src/config"
	"dynts-bann3r/src/image"
	"dynts-bann3r/src/label"
	"dynts-bann3r/src/teamspeak"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	client := teamspeak.Login(cfg.Connection)

	defer client.Close()

	for i, val := range cfg.Labels {
		text, err := label.GenerateLabel(val.Text, client)

		if err != nil {
			log.Fatal(err)
		}

		cfg.Labels[i].Text = text
	}

	image.AddLabelsToImage(cfg.Labels, "template.png", "banner.png")
}
