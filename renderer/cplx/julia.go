package cplx

func calcJulia(width uint, height uint, maxiter uint, palette [][3]byte, image []byte, zy0 float64, cx float64, cy float64, done chan bool) {
	var zx0 float64 = -2.0
	for x := uint(0); x < width; x++ {
		var zx float64 = zx0
		var zy float64 = zy0
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
		zx0 += 4.0 / float64(width)
	}
	done <- true
}

func RenderJuliaFractal(width uint, height uint, maxiter uint, palette [][3]byte) []byte {
	done := make(chan bool, height)

	image := make([]byte, width*height*3)
	offset := uint(0)
	delta := width * 3

	var zy0 float64 = -2.0
	for y := uint(0); y < height; y++ {
		go calcJulia(width, height, maxiter, palette, image[offset:offset+delta], zy0, 0.0, 1.0, done)
		offset += delta
		zy0 += 4.0 / float64(height)
	}
	for i := uint(0); i < height; i++ {
		println(i)
		<-done
	}
	return image
}
