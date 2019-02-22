package image

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func WritePNGImage(filename string, width uint, height uint, raw_data []byte) {
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	offset := uint(0)
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			var red byte = raw_data[offset]
			offset++
			var green byte = raw_data[offset]
			offset++
			var blue byte = raw_data[offset]
			offset++
			c := color.RGBA{red, green, blue, 255}
			img.SetRGBA(x, y, c)
		}
	}

	outfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}
