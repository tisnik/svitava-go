package cplx

func CalcBarnsley(width uint, height uint, maxiter uint, zimage_line []ZPixel, cy float64, done chan bool) {
	var cx float64 = -2.0
	for x := uint(0); x < width; x++ {
		var zx float64 = 0.0
		var zy float64 = 0.0
		var i uint = 0
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
		zimage_line[x] = ZPixel{Iter: i, Z: complex(zx, zy)}
		cx += 4.0 / float64(width)
	}
	done <- true
}
