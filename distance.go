package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/blas/blas32"
)

func cosineSimNative(a, b []float32) (float32, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("vectors have different dimensions")
	}

	var (
		sumProduct float64
		sumASquare float64
		sumBSquare float64
	)

	for i := range a {
		sumProduct += float64(a[i] * b[i])
		sumASquare += float64(a[i] * a[i])
		sumBSquare += float64(b[i] * b[i])
	}

	return float32(sumProduct / (math.Sqrt(sumASquare) * math.Sqrt(sumBSquare))), nil
}

func cosineSimBlas(a, b blas32.Vector) (float32, error) {
	if a.N != b.N {
		return 0, fmt.Errorf("vectors have different dimensions")
	}

	sumProduct := blas32.Dot(a, b)
	sumASquare := blas32.Dot(a, a)
	sumBSquare := blas32.Dot(b, b)

	return float32(float64(sumProduct) / (math.Sqrt(float64(sumASquare)) * math.Sqrt(float64(sumBSquare)))), nil
}

//go:noescape
func _cosineSimAsm(a []float32, b []float32) float32 // implemented externally

func cosineSimAsm(a, b []float32) (float32, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("vectors have different dimensions")
	}

	return _cosineSimAsm(a, b), nil
}
