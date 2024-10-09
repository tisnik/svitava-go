package cplx

// CalcBarnsleyJ1 calculates Barnsley J1 Mandelbrot-like set
func CalcBarnsleyJ1(
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
