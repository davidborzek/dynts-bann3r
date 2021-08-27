package image

import (
	"dynts-bann3r/src/config"
	"log"

	"github.com/fogleman/gg"
)

func AddLabelsToImage(labels []config.Label, imagePath string, outputName string) {
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
		if err := dc.LoadFontFace("fonts/"+label.Font, label.FontSize); err != nil {
			panic(err)
		}

		dc.DrawStringAnchored(label.Text, label.X, label.Y, 0.5, float64(gg.AlignCenter))
	}

	dc.Clip()
	dc.SavePNG(outputName)
}
