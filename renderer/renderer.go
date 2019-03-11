package renderer

import (
	"svitava/renderer/cplx"
	"image")

func init() {
	println("Init")
}

func RenderMandelbrotFractal(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	done := make(chan bool, height)

	zimage := cplx.NewZImage(width, height)

	var cy float64 = -1.5
	for y := uint(0); y < height; y++ {
		go cplx.CalcMandelbrot(width, height, maxiter, zimage[y], cy, done)
		cy += 3.0 / float64(height)
	}
	for i := uint(0); i < height; i++ {
		println(i)
		<-done
	}

	image := image.NewNRGBA(image.Rect(0, 0, int(width), int(height)))

        for y := 0; y < int(height); y++ {
		offset := image.PixOffset(0, y)
                println(y)
                for x := uint(0); x < width; x++ {
                        i := byte(zimage[y][x].Iter)
                        image.Pix[offset] = palette[i][0]
			offset++
                        image.Pix[offset] = palette[i][1]
			offset++
                        image.Pix[offset] = palette[i][2]
			offset++
                        image.Pix[offset] = 0xff
			offset++
                }
        }

	return image
}
