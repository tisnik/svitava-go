package cplx_test

import (
	"testing"

	"github.com/tisnik/svitava-go/renderer/cplx"
)

const (
	WIDTH   = 512
	HEIGHT  = 512
	MAXITER = 1000
)

func BenchmarkCalcMandelbrot(b *testing.B) {
	zimage := cplx.NewZImage(WIDTH, HEIGHT)
	for i := 0; i < b.N; i++ {
		cplx.CalcMandelbrot(WIDTH, HEIGHT, 0.0, 0.0, MAXITER, zimage)
	}
}

func BenchmarkCalcMandelbrotComplex(b *testing.B) {
	zimage := cplx.NewZImage(WIDTH, HEIGHT)
	for i := 0; i < b.N; i++ {
		cplx.CalcMandelbrotComplex(WIDTH, HEIGHT, 0.0, 0.0, MAXITER, zimage)
	}
}
