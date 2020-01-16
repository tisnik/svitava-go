//
//  (C) Copyright 2019, 2020  Pavel Tisnovsky
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
	"image"
	"image/png"
	"os"
)

// WritePNGImage writes an image represented by byte slice into file with PNG format.
func WritePNGImage(filename string, img image.Image) {
	outfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}
