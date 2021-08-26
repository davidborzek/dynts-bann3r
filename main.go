package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fogleman/gg"
	"github.com/spf13/viper"
)

type label struct {
	Text     string
	X        float64
	Y        float64
	FontSize float64
	Font     string
	Color    string
}

type config struct {
	Labels []label
}

func main() {
	cfg := loadConfig()
	addLabelsToImage(cfg.Labels, "banner.png", "out.png")
}

func addLabelsToImage(labels []label, imagePath string, outputName string) {
	image, err := gg.LoadImage(imagePath)

	if err != nil {
		log.Fatal(err)
	}

	imgWidth := image.Bounds().Dx()
	imgHeight := image.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(image, 0, 0)

	for _, label := range labels {
		dc.SetHexColor(label.Color)
		if err := dc.LoadFontFace(label.Font, label.FontSize); err != nil {
			panic(err)
		}

		dc.DrawStringAnchored(label.Text, label.X, label.Y, 0.5, float64(gg.AlignCenter))
	}

	dc.Clip()
	dc.SavePNG(outputName)
}

func loadConfig() config {
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	var C config

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.Unmarshal(&C)
	return C
}
