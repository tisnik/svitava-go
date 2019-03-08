package renderer

import "svitava/renderer/cplx"

func init() {
	println("Init")
}

func RenderMandelbrotFractal(width uint, height uint, maxiter uint, palette [][3]byte) []byte {
	done := make(chan bool, height)

        zimage := make([][]cplx.ZPixel, height)
        for i := uint(0); i<height; i++ {
                zimage[i] = make([]cplx.ZPixel, width)
        }

	var cy float64 = -1.5
	for y := uint(0); y < height; y++ {
		go cplx.CalcMandelbrot(width, height, maxiter, zimage[y], cy, done)
		cy += 3.0 / float64(height)
	}
	for i := uint(0); i < height; i++ {
		println(i)
		<-done
	}

	image := make([]byte, width*height*3)
	offset := uint(0)
	//delta := width * 3
        for y := uint(0); y < height; y++ {
                println(y)
                for x := uint(0); x < width; x++ {
                        i := byte(zimage[y][x].Iter)
                        image[offset] = palette[i][0]
                        image[offset+1] = palette[i][1]
                        image[offset+2] = palette[i][2]
                        offset += 3
                }
                //offset += delta
        }

	return image
}
