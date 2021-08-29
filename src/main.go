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
	"time"
)

var ts teamspeak.Teamspeak
var cfg config.Config

func main() {
	cfg = config.LoadConfig()
	ts = teamspeak.New(cfg.Connection, cfg.AdminGroups)

	port := os.Getenv("DYNTS_BANN3R_PORT")
	if port == "" {
		port = "9000"
	}

	log.Printf("[INFO] dynts-bann3r is starting - refreshing information every %d seconds \n", cfg.RefreshInterval)

	go serveBanner(port)
	scheduleRefresh()
}

func scheduleRefresh() {
	for {
		log.Println("[INFO] Refreshing teamspeak server information.")
		ts.Refresh()
		time.Sleep(time.Duration(cfg.RefreshInterval) * time.Second)
	}
}

func getBanner() i.Image {
	filledLabels := make([]config.Label, len(cfg.Labels))
	copy(filledLabels, cfg.Labels)

	for i, val := range cfg.Labels {
		filledLabels[i].Text = label.ReplacePlaceholders(val.Text, ts.State())
	}

	return image.AddLabelsToImage(filledLabels, cfg.TemplatePath)
}

func serveBanner(port string) {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		buffer := new(bytes.Buffer)
		err := png.Encode(buffer, getBanner())

		if err == nil {
			rw.Write(buffer.Bytes())
		}
	})

	log.Println("[http] Serving the banner on localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
