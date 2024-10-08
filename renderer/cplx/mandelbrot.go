package cplx

// CalcMandelbrot calculates Mandelbrot set into the provided ZPixels
func CalcMandelbrot(
	width uint, height uint, pcx float64, pcy float64,
	maxiter uint, zimage ZImage) {

	var cy float64 = -1.5
	for y := uint(0); y < height; y++ {
		var cx float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = pcx
			var zy float64 = pcy
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
			cx += 3.0 / float64(width)
		}
		cy += 3.0 / float64(height)
	}
}

// CalcMandelbrotComplex calculates Mandelbrot set into the provided ZPixels
// Calculations use complex numbers
func CalcMandelbrotComplex(
	width uint, height uint, pcx float64, pcy float64,
	maxiter uint, zimage ZImage) {

	var cy float64 = -1.5
	for y := uint(0); y < height; y++ {
		var cx float64 = -2.0
		for x := uint(0); x < width; x++ {
			var c complex128 = complex(cx, cy)
			var z complex128 = complex(pcx, pcy)
			var i uint
			for i < maxiter {
				zx := real(z)
				zy := imag(z)
				if zx*zx+zy*zy > 4.0 {
					break
				}
				z = z*z + c
				i++
			}
			zimage[y][x] = ZPixel{Iter: i, Z: z}
			cx += 3.0 / float64(width)
		}
		cy += 3.0 / float64(height)
	}
}
