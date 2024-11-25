package deepimage

import "image"

// IPixel is a representation of pixel as one unsighed integer value
type IPixel uint64

// IImage is representation of raster image consisting of IPixels
type IImage [][]IPixel

type DeepImage struct {
	Z    ZImage
	R    RImage
	I    IImage
	RGBB image.Image
}
