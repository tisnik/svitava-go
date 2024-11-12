//
//  (C) Copyright 2024  Pavel Tisnovsky
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
	"image/gif"
	"os"
)

// WriteGIFImage writes an image represented by standard image.Image structure into file with GIF format.
func WriteGIFImage(filename string, img image.Image) {
	outfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	gif.Encode(outfile, img, nil)
}
