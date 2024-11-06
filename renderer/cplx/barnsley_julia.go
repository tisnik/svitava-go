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

package cplx

import "github.com/tisnik/svitava-go/params"

// CalcBarnsleyJ1 calculates Barnsley J1 Mandelbrot-like set
func CalcBarnsleyJ1(
	width uint, height uint,
	params params.Cplx,
	zimage ZImage) {

	cx := params.Cx0
	cy := params.Cy0
	var zy0 float64 = -2.0
	for y := uint(0); y < height; y++ {
		var zx0 float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = zx0
			var zy float64 = zy0
			var i uint
			for i < params.Maxiter {
				var zxn float64
				var zyn float64
				zx2 := zx * zx
				zy2 := zy * zy
				if zx2+zy2 > 4.0 {
					break
				}
				if zx >= 0 {
					zxn = zx*cx - zy*cy - cx
					zyn = zx*cy + zy*cx - cy
				} else {
					zxn = zx*cx - zy*cy + cx
					zyn = zx*cy + zy*cx + cy
				}
				zx = zxn
				zy = zyn
				i++
			}
			zimage[y][x] = ZPixel{Iter: i, Z: complex(zx, zy)}
			zx0 += 4.0 / float64(width)
		}
		zy0 += 4.0 / float64(height)
	}
}

// CalcBarnsleyJ2 calculates Barnsley J2 Mandelbrot-like set
func CalcBarnsleyJ2(
	width uint, height uint,
	params params.Cplx,
	zimage ZImage) {

	cx := params.Cx0
	cy := params.Cy0
	var zy0 float64 = -2.0
	for y := uint(0); y < height; y++ {
		var zx0 float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = zx0
			var zy float64 = zy0
			var i uint
			for i < params.Maxiter {
				var zxn float64
				var zyn float64
				zx2 := zx * zx
				zy2 := zy * zy
				if zx2+zy2 > 4.0 {
					break
				}
				if zx*cy+zy*cx >= 0 {
					zxn = zx*cx - zy*cy - cx
					zyn = zx*cy + zy*cx - cy
				} else {
					zxn = zx*cx - zy*cy + cx
					zyn = zx*cy + zy*cx + cy
				}
				zx = zxn
				zy = zyn
				i++
			}
			zimage[y][x] = ZPixel{Iter: i, Z: complex(zx, zy)}
			zx0 += 4.0 / float64(width)
		}
		zy0 += 4.0 / float64(height)
	}
}

// CalcBarnsleyJ3 calculates Barnsley J3 Mandelbrot-like set
func CalcBarnsleyJ3(
	width uint, height uint,
	params params.Cplx,
	zimage ZImage) {

	cx := params.Cx0
	cy := params.Cy0
	var zy0 float64 = -2.0
	for y := uint(0); y < height; y++ {
		var zx0 float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = zx0
			var zy float64 = zy0
			var i uint
			for i < params.Maxiter {
				var zxn float64
				var zyn float64
				zx2 := zx * zx
				zy2 := zy * zy
				if zx2+zy2 > 4.0 {
					break
				}
				if zx > 0 {
					zxn = zx2 - zy2 - 1
					zyn = 2.0 * zx * zy
				} else {
					zxn = zx2 - zy2 - 1 + cx*zx
					zyn = 2.0*zx*zy + cy*zx
				}
				zx = zxn
				zy = zyn
				i++
			}
			zimage[y][x] = ZPixel{Iter: i, Z: complex(zx, zy)}
			zx0 += 4.0 / float64(width)
		}
		zy0 += 4.0 / float64(height)
	}
}
