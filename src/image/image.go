package image

import (
	"dynts-bann3r/src/config"
	"image"
	"log"

	"github.com/fogleman/gg"
)

func AddLabelsToImage(labels []config.Label, imagePath string) image.Image {
	image, err := gg.LoadImage(imagePath)

	if err != nil {
		log.Fatal(err)
	}

	imgWidth := image.Bounds().Dx()
	imgHeight := image.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(image, 0, 0)

	for _, label := range labels {

		if label.Color == "" {
			label.Color = "#000000"
		}

		if label.Font == "" {
			label.Font = "Arial.ttf"
		}

		if label.FontSize == 0 {
			label.FontSize = 16
		}

		dc.SetHexColor(label.Color)
		if err := dc.LoadFontFace("fonts/"+label.Font, label.FontSize); err != nil {
			log.Println("[WARN] font '" + label.Font + "' could not be loaded - using default font 'Arial.ttf'")

			if err := dc.LoadFontFace("fonts/Arial.ttf", label.FontSize); err != nil {
				log.Printf("[ERROR] font Arial.ttf could not be loaded: %v \n", err)
			}
		}

		dc.DrawStringAnchored(label.Text, label.X, label.Y, 0.5, float64(gg.AlignCenter))
	}

	dc.Clip()
	return dc.Image()
}
