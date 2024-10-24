//
//  (C) Copyright 2019, 2020, 2021, 2022, 2023, 2024  Pavel Tisnovsky
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

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tisnik/svitava-go/configuration"
	"github.com/tisnik/svitava-go/image"
	"github.com/tisnik/svitava-go/server"

	//"github.com/tisnik/svitava/palettes"
	"github.com/tisnik/svitava-go/params"
	"github.com/tisnik/svitava-go/renderer"
)

const (
	CONFIG_FILE_NAME = "config.toml"
)

func main() {
	var width uint
	var height uint
	var aa bool
	var startServer bool
	var startTUI bool
	var execute string
	var port uint

	configuration, err := configuration.LoadConfiguration(CONFIG_FILE_NAME)
	if err != nil {
		println("Unable to load configuration")
		os.Exit(1)
	}
	fmt.Println(configuration)

	flag.UintVar(&width, "w", 0, "image width (shorthand)")
	flag.UintVar(&width, "width", 0, "image width")

	flag.UintVar(&height, "h", 0, "image height (shorthand)")
	flag.UintVar(&height, "height", 0, "image height")

	flag.BoolVar(&aa, "a", false, "enable antialiasing (shorthand)")
	flag.BoolVar(&aa, "antialias", false, "enable antialiasing")

	flag.BoolVar(&startTUI, "t", false, "start with text user interface (shorthand)")
	flag.BoolVar(&startTUI, "tui", false, "start in server mode")

	flag.StringVar(&execute, "e", "", "execute given script with rendering commands (shorthand)")
	flag.StringVar(&execute, "exec", "", "execute given script with rendering commands")
	flag.StringVar(&execute, "execute", "", "execute given script with rendering commands")

	flag.BoolVar(&startServer, "s", false, "start in server mode (shorthand)")
	flag.BoolVar(&startServer, "server", false, "start in server mode")

	flag.UintVar(&port, "p", 8080, "port for the server (shorthand)")
	flag.UintVar(&port, "port", 8080, "port for the server")

	flag.Parse()

	if startServer {
		log.Println("Starting server")
		server.StartServer(port)
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
		image.WritePNGImage("mandelbrot.png", img)

		img = renderer.RenderBarnsleyFractalM1(width, height, 255, palette[:])
		image.WritePNGImage("barnsley_m1.png", img)

		img = renderer.RenderBarnsleyFractalJ1(width, height, 255, palette[:])
		image.WritePNGImage("barnsley_j1.png", img)

		img = renderer.RenderBarnsleyFractalM2(width, height, 255, palette[:])
		image.WritePNGImage("barnsley_m2.png", img)

		img = renderer.RenderBarnsleyFractalJ2(width, height, 255, palette[:])
		image.WritePNGImage("barnsley_j2.png", img)

		img = renderer.RenderBarnsleyFractalM3(width, height, 255, palette[:])
		image.WritePNGImage("barnsley_m3.png", img)

		img = renderer.RenderBarnsleyFractalJ3(width, height, 255, palette[:])
		image.WritePNGImage("barnsley_j3.png", img)

		img2 := renderer.RenderJuliaFractal(width, height, 255, palette[:])
		image.WritePNGImage("julia.png", img2)

		img2 = renderer.RenderMagnetFractal(width, height, 255, palette[:])
		image.WritePNGImage("magnet.png", img2)

		//img = renderer.RenderBarnsleyFractalJ1(width, height, 255, palette[:])
		//image.WritePNGImage("barnsley_j1.png", img)
	}

	parameters, err := params.LoadCplxParameters("data/complex_fractals.toml")
	fmt.Printf("%v\n%v\n", parameters, err)
}
