package renderer

import (
	"image"
	"svitava-go/renderer/cplx"
)

func init() {
	println("Init")
}

func waitForWorkers(done chan bool, height uint) {
	for i := uint(0); i < height; i++ {
		println(i)
		<-done
	}
}

func complexImageToImage(zimage cplx.ZImage, width uint, height uint, palette [][3]byte) image.Image {
	image := image.NewNRGBA(image.Rect(0, 0, int(width), int(height)))

	for y := 0; y < int(height); y++ {
		offset := image.PixOffset(0, y)
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

// RenderMandelbrotFractal renders a classic Mandelbrot fractal into provided Image.
func RenderMandelbrotFractal(width uint, height uint, pcx float64, pcy float64, maxiter uint, palette [][3]byte) image.Image {
	done := make(chan bool, height)

	zimage := cplx.NewZImage(width, height)

	var cy float64 = -1.5
	for y := uint(0); y < height; y++ {
		go cplx.CalcMandelbrot(width, height, pcx, pcy, maxiter, zimage[y], cy, done)
		cy += 3.0 / float64(height)
	}
	waitForWorkers(done, height)

	return complexImageToImage(zimage, width, height, palette)
}

// RenderJuliaFractal renders a classic Julia fractal into provided Image.
func RenderJuliaFractal(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	done := make(chan bool, height)

	zimage := cplx.NewZImage(width, height)

	var zy0 float64 = -2.0
	for y := uint(0); y < height; y++ {
		go cplx.CalcJulia(width, height, maxiter, zimage[y], zy0, 0.0, 1.0, done)
		zy0 += 4.0 / float64(height)
	}
	for i := uint(0); i < height; i++ {
		println(i)
		<-done
	}

	return complexImageToImage(zimage, width, height, palette)
}

// RenderBarnsleyFractalM1 renders a classic Barnsley fractal M1 into provided Image.
func RenderBarnsleyFractalM1(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	done := make(chan bool, height)

	zimage := cplx.NewZImage(width, height)

	var cy float64 = -2.0
	for y := uint(0); y < height; y++ {
		go cplx.CalcBarnsleyM1(width, height, maxiter, zimage[y], cy, done)
		cy += 4.0 / float64(height)
	}
	waitForWorkers(done, height)

	return complexImageToImage(zimage, width, height, palette)
}

// RenderBarnsleyFractalM2 renders a classic Barnsley fractal M2 into provided Image.
func RenderBarnsleyFractalM2(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	done := make(chan bool, height)

	zimage := cplx.NewZImage(width, height)

	var cy float64 = -2.0
	for y := uint(0); y < height; y++ {
		go cplx.CalcBarnsleyM2(width, height, maxiter, zimage[y], cy, done)
		cy += 4.0 / float64(height)
	}
	waitForWorkers(done, height)

	return complexImageToImage(zimage, width, height, palette)
}

// RenderBarnsleyFractalM3 renders a classic Barnsley fractal M3 into provided Image.
func RenderBarnsleyFractalM3(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	done := make(chan bool, height)

	zimage := cplx.NewZImage(width, height)

	var cy float64 = -2.0
	for y := uint(0); y < height; y++ {
		go cplx.CalcBarnsleyM3(width, height, maxiter, zimage[y], cy, done)
		cy += 4.0 / float64(height)
	}
	waitForWorkers(done, height)

	return complexImageToImage(zimage, width, height, palette)
}
