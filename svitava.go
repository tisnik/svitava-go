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
	"log"
	//"svitava-go/server"
	"svitava-go/image"
	//"svitava/palettes"
	"svitava-go/renderer"
	//"svitava-go/renderer/cplx"
)

func main() {
	var width uint
	var height uint
	var aa bool
	var startServer bool
	var port uint

	flag.UintVar(&width, "w", 0, "image width (shorthand)")
	flag.UintVar(&width, "width", 0, "image width")

	flag.UintVar(&height, "h", 0, "image height (shorthand)")
	flag.UintVar(&height, "height", 0, "image height")

	flag.BoolVar(&aa, "a", false, "enable antialiasing (shorthand)")
	flag.BoolVar(&aa, "antialias", false, "enable antialiasing")

	flag.BoolVar(&startServer, "s", false, "start in server mode (shorthand)")
	flag.BoolVar(&startServer, "server", false, "start in server mode")

	flag.UintVar(&port, "p", 8080, "port for the server (shorthand)")
	flag.UintVar(&port, "port", 8080, "port for the server")

	flag.Parse()

	if startServer {
		log.Println("Starting server")
		//server.StartServer(port)
	} else {
		log.Println("Starting renderer")
		width = 256
		height = 256

		// palette := palettes.Mandmap
		palette := [256][3]byte{}
		for i := 0; i < 256; i++ {
			index := byte(i)
			palette[index][0] = index * 2
			palette[index][1] = index * 5
			palette[index][2] = index * 7
		}
		img := renderer.RenderMandelbrotFractal(width, height, 0.0, 0.0, 255, palette[:])
		image.WritePNGImage("mandelbror.png", img)

		// img = renderer.RenderBarnsleyFractalM1(width, height, 255, palette[:])
		// image.WritePNGImage("barnsley_m1.png", img)

		// img = renderer.RenderBarnsleyFractalM2(width, height, 255, palette[:])
		// image.WritePNGImage("barnsley_m2.png", img)

		// img = renderer.RenderBarnsleyFractalM3(width, height, 255, palette[:])
		// image.WritePNGImage("barnsley_m3.png", img)

		img2 := renderer.RenderJuliaFractal(width, height, 255, palette[:])
		image.WritePNGImage("julia.png", img2)

		// img = renderer.RenderBarnsleyFractalJ1(width, height, 255, palette[:])
		// image.WritePNGImage("barnsley_j1.png", img)
	}
}
