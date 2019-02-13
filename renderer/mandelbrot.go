package renderer

func calcMandelbrot(width uint, height uint, maxiter uint, palette [][3]byte, image []byte, cy float64) {
	var cx float64 = -2.0
	for x := uint(0); x < width; x++ {
		var zx float64 = 0.0
		var zy float64 = 0.0
		var i uint = 0
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
		color := palette[i]
		image[3*x] = color[0]
		image[3*x+1] = color[1]
		image[3*x+2] = color[2]
		cx += 3.0 / float64(width)
	}
}
