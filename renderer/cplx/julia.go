package cplx

import "github.com/tisnik/svitava-go/params"

// CalcJulia calculates classic Julia fractal
func CalcJulia(
	width uint, height uint,
	params params.Cplx,
	zimage ZImage) {

	var zy0 float64 = -2.0
	for y := uint(0); y < height; y++ {
		var zx0 float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = zx0
			var zy float64 = zy0
			var i uint
			for i < params.Maxiter {
				zx2 := zx * zx
				zy2 := zy * zy
				if zx2+zy2 > 4.0 {
					break
				}
				zy = 2.0*zx*zy + params.Cy0
				zx = zx2 - zy2 + params.Cx0
				i++
			}
			zimage[y][x] = ZPixel{Iter: i, Z: complex(zx, zy)}
			zx0 += 4.0 / float64(width)
		}
		zy0 += 4.0 / float64(height)
	}
}
