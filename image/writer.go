package image

import "image"

type Writer interface {
	WriteImage(filename string, img image.Image) error
}
