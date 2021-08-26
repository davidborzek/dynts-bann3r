package main

import (
	"dynts-bann3r/src/config"
	"dynts-bann3r/src/image"
)

func main() {
	cfg := config.LoadConfig()
	image.AddLabelsToImage(cfg.Labels, "template.png", "banner.png")
}
