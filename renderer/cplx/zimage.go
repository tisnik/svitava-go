package cplx

type ZImage [][]ZPixel

func NewZImage(width uint, height uint) ZImage {
        zimage := make([][]ZPixel, height)
        for i := uint(0); i<height; i++ {
                zimage[i] = make([]ZPixel, width)
        }
	return zimage
}
