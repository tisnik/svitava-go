package cplx

// CalcJulia calculates classic Julia fractal
func CalcJulia(
	width uint, height uint,
	cx float64, cy float64,
	maxiter uint, zimage ZImage) {

	var zy0 float64 = -2.0
	for y := uint(0); y < height; y++ {
		var zx0 float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = zx0
			var zy float64 = zy0
			var i uint
			for i < maxiter {
				zx2 := zx * zx
				zy2 := zy * zy
				if zx2+zy2 > 4.0 {
					break
				}
				zy = 2.0*zx*zy + cy
				zx = zx2 - zy2 + cx
				i++
			}
			zimage[y][x] = ZPixel{Iter: i, Z: complex(zx, zy)}
			zx0 += 4.0 / float64(width)
		}
		zy0 += 4.0 / float64(height)
	}
}
