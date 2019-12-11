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

package image

import (
	"bufio"
	"fmt"
	"os"
)

// WritePPMImage writes an image represented by byte slice into file with PPM format.
func WritePPMImage(width uint, height uint, image []byte) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprintln(w, "P3")
	fmt.Fprintf(w, "%d %d\n", width, height)
	fmt.Fprintln(w, "255")

	for i := 0; i < len(image); {
		r := image[i]
		i++
		g := image[i]
		i++
		b := image[i]
		i++
		fmt.Fprintf(w, "%d %d %d\n", r, g, b)
	}
}
