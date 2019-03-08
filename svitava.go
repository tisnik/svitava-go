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

package main

// import "svitava/renderer"
import (
	"flag"
	"svitava/image"
	//"svitava/palettes"
	"svitava/renderer"
	//"svitava/renderer/cplx"
)

func main() {
	var width uint
	var height uint
	var aa bool

	flag.UintVar(&width, "w", 0, "image width (shorthand)")
	flag.UintVar(&width, "width", 0, "image width")

	flag.UintVar(&height, "h", 0, "image height (shorthand)")
	flag.UintVar(&height, "height", 0, "image height")

	flag.BoolVar(&aa, "a", false, "enable antialiasing (shorthand)")
	flag.BoolVar(&aa, "antialias", false, "enable antialiasing")

	flag.Parse()

	width = 256
	height = 256

	// palette := palettes.Mandmap
	palette := [256][3]byte{}
        for i := 0; i < 256; i++ {
                index := byte(i)
                palette[index][0] = index * 2
                palette[index][1] = index * 3
                palette[index][2] = index * 4
        }
	img := renderer.RenderMandelbrotFractal(width, height, 255, palette[:])
	// image.WritePPMImage(width, height, img)
	image.WritePNGImage("mandel.png", width, height, img)

	// img2 := cplx.RenderJuliaFractal(width, height, 255, palette[:])
	// image.WritePNGImage("julia.png", width, height, img2)
}
