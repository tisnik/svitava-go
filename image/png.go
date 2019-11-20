package image

import (
	"image"
	"image/png"
	"os"
)

// WritePNGImage writes an image represented by byte slice into file with PNG format.
func WritePNGImage(filename string, img image.Image) {
	outfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}
