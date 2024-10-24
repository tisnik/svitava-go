package cplx

import "github.com/tisnik/svitava-go/params"

// CalcBarnsleyM1 calculates Barnsley M1 Mandelbrot-like set
func CalcBarnsleyM1(
	width uint, height uint,
	params params.Cplx,
	zimage ZImage) {

	var cy float64 = -2.0
	for y := uint(0); y < height; y++ {
		var cx float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = cx
			var zy float64 = cy
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
			cx += 4.0 / float64(width)
		}
		cy += 4.0 / float64(height)
	}
}

// CalcBarnsleyM2 calculates Barnsley M2 Mandelbrot-like set
func CalcBarnsleyM2(
	width uint, height uint,
	params params.Cplx,
	zimage ZImage) {

	var cy float64 = -2.0
	for y := uint(0); y < height; y++ {
		var cx float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = cx
			var zy float64 = cy
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
			cx += 4.0 / float64(width)
		}
		cy += 4.0 / float64(height)
	}
}

// CalcBarnsleyM3 calculates Barnsley M3 Mandelbrot-like set
func CalcBarnsleyM3(
	width uint, height uint,
	params params.Cplx,
	zimage ZImage) {

	var cy float64 = -2.0
	for y := uint(0); y < height; y++ {
		var cx float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = cx
			var zy float64 = cy
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
			cx += 4.0 / float64(width)
		}
		cy += 4.0 / float64(height)
	}
}
