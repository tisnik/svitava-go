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
)

func main() {
	var width int
	var height int
	var aa bool

	flag.IntVar(&width, "w", 0, "image width (shorthand)")
	flag.IntVar(&width, "width", 0, "image width")

	flag.IntVar(&height, "h", 0, "image height (shorthand)")
	flag.IntVar(&height, "height", 0, "image height")

	flag.BoolVar(&aa, "a", false, "enable antialiasing (shorthand)")
	flag.BoolVar(&aa, "antialias", false, "enable antialiasing")

	flag.Parse()
}
