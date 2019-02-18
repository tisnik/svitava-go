package renderer

func calcMandelbrot(width uint, height uint, maxiter uint, palette [][3]byte, image []byte, cy float64, done chan bool) {
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
	done <- true
}

func calcMandelbrotComplex(width uint, height uint, maxiter uint, palette [][3]byte, image []byte, cy float64, done chan bool) {
	var c complex128 = complex(-2.0, cy)
	var dc complex128 = complex(3.0/float64(width), 0.0)
	for x := uint(0); x < width; x++ {
		var z complex128 = 0.0 + 0.0i
		var i uint = 0
		for i < maxiter {
			zx := real(z)
			zy := imag(z)
			if zx*zx+zy*zy > 4.0 {
				break
			}
			z = z*z + c
			i++
		}
		color := palette[i]
		image[3*x] = color[0]
		image[3*x+1] = color[1]
		image[3*x+2] = color[2]
		c += dc
	}
	done <- true
}

func RenderMandelbrotFractal(width uint, height uint, maxiter uint, palette [][3]byte) []byte {
	done := make(chan bool, height)

	image := make([]byte, width*height*3)
	offset := uint(0)
	delta := width * 3

	var cy float64 = -1.5
	for y := uint(0); y < height; y++ {
		go calcMandelbrot(width, height, maxiter, palette, image[offset:offset+delta], cy, done)
		offset += delta
		cy += 3.0 / float64(height)
	}
	for i := uint(0); i < height; i++ {
		println(i)
		<-done
	}
	return image
}
