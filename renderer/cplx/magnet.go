package cplx

// CalcMagnet calculates Magnet Mandelbrot-like set
func CalcMagnet(
	width uint, height uint,
	pcx float64, pcy float64,
	maxiter uint, zimage ZImage) {

	var cy float64 = -2.0
	for y := uint(0); y < height; y++ {
		var cx float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = cx
			var zy float64 = cy
			var i uint
			for i < maxiter {
				var zxn float64
				var zyn float64
				zx2 := zx * zx
				zy2 := zy * zy
				if zx2+zy2 > 100.0 {
					break
				}
				if ((zx-1.0)*(zx-1.0) + zy*zy) < 0.001 {
					break
				}
				tzx := zx2 - zy2 + cx - 1
				tzy := 2.0*zx*zy + cy
				bzx := 2.0*zx + cx - 2
				bzy := 2.0*zy + cy
				div := bzx*bzx + bzy*bzy
				const MIN_VALUE = 1.0 - 100
				if div < MIN_VALUE {
					break
				}
				zxn = (tzx*bzx + tzy*bzy) / div
				zyn = (tzy*bzx - tzx*bzy) / div
				zx = (zxn + zyn) * (zxn - zyn)
				zy = 2.0 * zxn * zyn
				i++
			}
			zimage[y][x] = ZPixel{Iter: i, Z: complex(zx, zy)}
			cx += 4.0 / float64(width)
		}
		cy += 4.0 / float64(height)
	}
}
