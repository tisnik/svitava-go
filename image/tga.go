//
//  (C) Copyright 2024  Pavel Tisnovsky
//
//  All rights reserved. This program and the accompanying materials
//  are made available under the terms of the Eclipse Public License v1.0
//  which accompanies this distribution, and is available at
//  http://www.eclipse.org/legal/epl-v10.html
//
//  Contributors:
//      Pavel Tisnovsky
//

package image

import "image"

// TGAImageWriter implements image.Writer interface, it writes TGA format
type TGAImageWriter struct{}

// WriteBMPImage writes an image represented by byte slice into file with TGA format.
func (writer TGAImageWriter) WriteImage(filename string, img image.Image) error {
	return nil
}

// NewTGAImageWriter is a constructor for TGA image writer
func NewTGAImageWriter() TGAImageWriter {
	return TGAImageWriter{}
}
