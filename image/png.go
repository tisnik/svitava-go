package image

import (
	"image"
	"image/png"
	"os"
)

func WritePNGImage(filename string, img image.Image) {
	outfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}
