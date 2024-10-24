package cplx_test

import (
	"testing"

	"github.com/tisnik/svitava-go/params"
	"github.com/tisnik/svitava-go/renderer/cplx"
)

const (
	WIDTH   = 512
	HEIGHT  = 512
	MAXITER = 1000
)

func BenchmarkCalcMandelbrot(b *testing.B) {
	params := params.Cplx{
		Cx0:     0,
		Cy0:     0,
		Maxiter: 1000,
	}
	zimage := cplx.NewZImage(WIDTH, HEIGHT)
	for i := 0; i < b.N; i++ {
		cplx.CalcMandelbrot(WIDTH, HEIGHT, params, zimage)
	}
}

func BenchmarkCalcMandelbrotComplex(b *testing.B) {
	params := params.Cplx{
		Cx0:     0,
		Cy0:     0,
		Maxiter: 1000,
	}
	zimage := cplx.NewZImage(WIDTH, HEIGHT)
	for i := 0; i < b.N; i++ {
		cplx.CalcMandelbrotComplex(WIDTH, HEIGHT, params, zimage)
	}
}

func BenchmarkCalcJulia(b *testing.B) {
	zimage := cplx.NewZImage(WIDTH, HEIGHT)
	for i := 0; i < b.N; i++ {
		cplx.CalcJulia(WIDTH, HEIGHT, 0.0, 1.0, MAXITER, zimage)
	}
}
