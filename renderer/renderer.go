//
//  (C) Copyright 2019 - 2025  Pavel Tisnovsky
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

	"github.com/tisnik/svitava-go/deepimage"
	im "github.com/tisnik/svitava-go/image"
	"github.com/tisnik/svitava-go/palettes"
	"github.com/tisnik/svitava-go/params"
	"github.com/tisnik/svitava-go/renderer/cplx"
	"github.com/tisnik/svitava-go/renderer/textures"
)

func init() {
	println("Init")
}

type Renderer interface {
	RenderComplexFractal(resolution im.Resolution, params params.Cplx, palette palettes.Palette) image.Image
}

type SingleGoroutineRenderer struct {
}

func NewSingleGoroutineRenderer() Renderer {
	return SingleGoroutineRenderer{}
}

type fractalFunction = func(params params.Cplx, deepImage deepimage.Image)

func render(width uint, height uint, params params.Cplx, palette palettes.Palette, function fractalFunction) image.Image {
	if width == 0 || height == 0 {
		// TODO: logging
		return nil
	}
	if function == nil {
		// TODO: logging
		return nil
	}
	deepImage, err := deepimage.New(width, height)
	if err != nil {
		// TODO: logging
		return nil
	}
	function(params, deepImage)
	deepImage.ApplyPalette(palette)
	return deepImage.RGBA
}

func (r SingleGoroutineRenderer) RenderComplexFractal(
	resolution im.Resolution,
	params params.Cplx,
	palette palettes.Palette) image.Image {

	functions := map[string]fractalFunction{
		"Classic Mandelbrot set":             cplx.CalcMandelbrotComplex,
		"Classic Julia set":                  cplx.CalcJulia,
		"Mandelbrot set z=z^3+c":             cplx.CalcMandelbrotZ3,
		"Mandelbrot set z=z^4+c":             cplx.CalcMandelbrotZ4,
		"Mandelbrot set z=z^2+z+c":           cplx.CalcMandelbrotZ2pZ,
		"Mandelbrot set z=z^2-z+c":           cplx.CalcMandelbrotZ2mZ,
		"Mandelbrot set z=sin(z)*c":          cplx.CalcMandelbrotFn,
		"Julia set z=z^3+c":                  cplx.CalcJuliaZ3,
		"Julia set z=z^4+c":                  cplx.CalcJuliaZ4,
		"Julia set z=sin(z)*c":               cplx.CalcJuliaFn,
		"Barnsley set 1, Mandelbrot variant": cplx.CalcBarnsleyMandelbrotM1,
		"Barnsley set 1, Julia variant":      cplx.CalcBarnsleyJuliaJ1,
		"Barnsley set 2, Mandelbrot variant": cplx.CalcBarnsleyMandelbrotM2,
		"Barnsley set 2, Julia variant":      cplx.CalcBarnsleyJuliaJ2,
		"Barnsley set 3, Mandelbrot variant": cplx.CalcBarnsleyMandelbrotM3,
		"Barnsley set 3, Julia variant":      cplx.CalcBarnsleyJuliaJ3,
		"Phoenix set, Mandelbrot variant":    cplx.CalcPhoenixM,
		"Phoenix set, Julia variant":         cplx.CalcPhoenixJ,
		"Lambda, Mandelbrot variant":         cplx.CalcMandelLambda,
		"Lambda, Julia variant":              cplx.CalcLambda,
		"Manowar, Mandelbrot variant":        cplx.CalcManowarM,
		"Manowar, Julia variant":             cplx.CalcManowarJ,
		"Mandelbrot, Z power":                cplx.CalcZPowerMandelbrot,
		"Circle pattern":                     textures.CalcCirclePattern,
		"Plasma pattern":                     textures.CalcPlasmaPattern,
		"FM synth":                           textures.CalcFMSynth,
	}

	function, exists := functions[params.Name]
	if !exists {
		// Return default function
		function = cplx.CalcMandelbrotComplex
	}

	return render(resolution.Width, resolution.Height, params, palette, function)
}

// RenderBarnsleyFractalM2 renders a classic Barnsley fractal M2 into provided Image.
func RenderBarnsleyFractalM2(width uint, height uint, maxiter uint, palette palettes.Palette) image.Image {
	params := params.Cplx{
		Cx0:     0.0,
		Cy0:     1.0,
		Maxiter: maxiter,
	}
	return render(width, height, params, palette, cplx.CalcBarnsleyMandelbrotM2)
}

// RenderBarnsleyFractalM3 renders a classic Barnsley fractal M3 into provided Image.
func RenderBarnsleyFractalM3(width uint, height uint, maxiter uint, palette palettes.Palette) image.Image {
	params := params.Cplx{
		Cx0:     0.0,
		Cy0:     1.0,
		Maxiter: maxiter,
	}
	return render(width, height, params, palette, cplx.CalcBarnsleyMandelbrotM3)
}

// RenderBarnsleyFractalJ1 renders a classic Barnsley fractal J1 into provided Image.
func RenderBarnsleyFractalJ1(width uint, height uint, maxiter uint, palette palettes.Palette) image.Image {
	params := params.Cplx{
		Cx0:     0.48,
		Cy0:     -1.32,
		Maxiter: maxiter,
	}
	return render(width, height, params, palette, cplx.CalcBarnsleyJuliaJ1)
}

// RenderBarnsleyFractalJ2 renders a classic Barnsley fractal J2 into provided Image.
func RenderBarnsleyFractalJ2(width uint, height uint, maxiter uint, palette palettes.Palette) image.Image {
	params := params.Cplx{
		Cx0:     0.5,
		Cy0:     1.2,
		Maxiter: maxiter,
	}
	return render(width, height, params, palette, cplx.CalcBarnsleyJuliaJ2)
}

// RenderBarnsleyFractalJ3 renders a classic Barnsley fractal J3 into provided Image.
func RenderBarnsleyFractalJ3(width uint, height uint, maxiter uint, palette palettes.Palette) image.Image {
	params := params.Cplx{
		Cx0:     0.0,
		Cy0:     0.0,
		Maxiter: maxiter,
	}
	return render(width, height, params, palette, cplx.CalcBarnsleyJuliaJ3)
}

// RenderMagnet renders a classic Magnet fractal into provided Image.
func RenderMagnetFractal(width uint, height uint, maxiter uint, palette palettes.Palette) image.Image {
	/*params := params.Cplx{
		Cx0:     1.1,
		Cy0:     1.1,
		Maxiter: maxiter,
	}*/
	params := params.Cplx{
		Cx0:     0.0,
		Cy0:     0.0,
		Maxiter: maxiter,
	}
	return render(width, height, params, palette, cplx.CalcMagnet)
}

// RenderMagnet renders a classic Magnet Julia fractal into provided Image.
func RenderMagnetJuliaFractal(width uint, height uint, maxiter uint, palette palettes.Palette) image.Image {
	params := params.Cplx{
		Cx0:     0.5,
		Cy0:     -1.5,
		Maxiter: maxiter,
	}
	return render(width, height, params, palette, cplx.CalcMagnetJulia)
}
