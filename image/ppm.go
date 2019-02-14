package image

import (
	"bufio"
	"fmt"
	"os"
)

func writePPMImage(width uint, height uint, image []byte) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprintln(w, "P3")
	fmt.Fprintf(w, "%d %d\n", width, height)
	fmt.Fprintln(w, "255")

	for i := 0; i < len(image); {
		r := image[i]
		i++
		g := image[i]
		i++
		b := image[i]
		i++
		fmt.Fprintf(w, "%d %d %d\n", r, g, b)
	}
}
