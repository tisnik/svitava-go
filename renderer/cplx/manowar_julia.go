//
//  (C) Copyright 2025  Pavel Tisnovsky
//
//  All rights reserved. This program and the accompanying materials
//  are made available under the terms of the Eclipse Public License v1.0
//  which accompanies this distribution, and is available at
//  http://www.eclipse.org/legal/epl-v10.html
//
//  Contributors:
//      Pavel Tisnovsky
//

package cplx

import (
	"github.com/tisnik/svitava-go/deepimage"
	"github.com/tisnik/svitava-go/params"
)

// CalcManowarJ calculates Manowar Julia-like set
func CalcManowarJ(
	params params.Cplx,
	image deepimage.Image) {

	cx := params.Cx0
	cy := params.Cy0
	var c complex128 = complex(cx, cy)

	println(c)

	var zy0 float64 = 1.0
	for y := uint(0); y < image.Resolution.Height; y++ {
		var zx0 float64 = -1.5
		for x := uint(0); x < image.Resolution.Width; x++ {
			var z complex128 = complex(zx0, zy0)
			var z1 = z
			var i uint
			for i < params.Maxiter {
				zx := real(z)
				zy := imag(z)
				if zx*zx+zy*zy > 4.0 {
					break
				}
				z2 := z*z + z1 + c
				z1 = z
				z = z2
				i++
			}
			i *= 3
			image.Z[y][x] = deepimage.ZPixel(z)
			image.I[y][x] = deepimage.IPixel(i)
			zx0 += 2.0 / float64(image.Resolution.Width)
		}
		zy0 += -2.0 / float64(image.Resolution.Height)
	}
}
