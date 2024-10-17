//
//  (C) Copyright 2019  Pavel Tisnovsky
//
//  All rights reserved. This program and the accompanying materials
//  are made available under the terms of the Eclipse Public License v1.0
//  which accompanies this distribution, and is available at
//  http://www.eclipse.org/legal/epl-v10.html
//
//  Contributors:
//      Pavel Tisnovsky
//

package renderer

import (
	"image"

	"github.com/tisnik/svitava-go/params"
	"github.com/tisnik/svitava-go/renderer/cplx"
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

type fractalFunction = func(width uint, height uint, pcx float64, pcy float64, maxiter uint, zimage cplx.ZImage)
type fractalFunction2 = func(width uint, height uint, params params.Cplx, zimage cplx.ZImage)

func render(width uint, height uint, pcx float64, pcy float64, maxiter uint, palette [][3]byte, function fractalFunction) image.Image {
	zimage := cplx.NewZImage(width, height)
	function(width, height, pcx, pcy, maxiter, zimage)
	return complexImageToImage(zimage, width, height, palette)
}

func render2(width uint, height uint, params params.Cplx, palette [][3]byte, function fractalFunction2) image.Image {
	zimage := cplx.NewZImage(width, height)
	function(width, height, params, zimage)
	return complexImageToImage(zimage, width, height, palette)
}

// RenderMandelbrotFractal renders a classic Mandelbrot fractal into provided Image.
func RenderMandelbrotFractal(width uint, height uint, pcx float64, pcy float64, maxiter uint, palette [][3]byte) image.Image {
	params := params.Cplx{
		Cx0:     0,
		Cy0:     0,
		Maxiter: 1000,
	}
	return render2(width, height, params, palette, cplx.CalcMandelbrotComplex)
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
	return render(width, height, 0.0, 0.0, maxiter, palette, cplx.CalcBarnsleyM1)
}

// RenderBarnsleyFractalM2 renders a classic Barnsley fractal M2 into provided Image.
func RenderBarnsleyFractalM2(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	return render(width, height, 0.0, 0.0, maxiter, palette, cplx.CalcBarnsleyM2)
}

// RenderBarnsleyFractalM3 renders a classic Barnsley fractal M3 into provided Image.
func RenderBarnsleyFractalM3(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	return render(width, height, 0.0, 0.0, maxiter, palette, cplx.CalcBarnsleyM3)
}

// RenderBarnsleyFractalJ1 renders a classic Barnsley fractal J1 into provided Image.
func RenderBarnsleyFractalJ1(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	return render(width, height, 1.0, 1.1, maxiter, palette, cplx.CalcBarnsleyJ1)
}

// RenderBarnsleyFractalJ2 renders a classic Barnsley fractal J2 into provided Image.
func RenderBarnsleyFractalJ2(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	return render(width, height, 1.1, 1.1, maxiter, palette, cplx.CalcBarnsleyJ2)
}

// RenderBarnsleyFractalJ3 renders a classic Barnsley fractal J3 into provided Image.
func RenderBarnsleyFractalJ3(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	return render(width, height, 0.0, 0.0, maxiter, palette, cplx.CalcBarnsleyJ3)
}

// RenderMagnet renders a classic Magnet fractal into provided Image.
func RenderMagnetFractal(width uint, height uint, maxiter uint, palette [][3]byte) image.Image {
	return render(width, height, 1.1, 1.1, maxiter, palette, cplx.CalcMagnet)
}
