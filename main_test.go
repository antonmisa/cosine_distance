package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/blas/blas32"
)

var precision uint = 7
var exp float32 = 0.0000001
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func randFloats(min, max float32, n int) []float32 {
	res := make([]float32, n)
	for i := range res {
		res[i] = min + rnd.Float32()*(max-min)
	}
	return res
}

func roundFloat(val float32, precision uint) float32 {
	ratio := math.Pow(10, float64(precision))
	return float32(math.Round(float64(val)*ratio) / ratio)
}

var table = []struct {
	input int
	data  [][]float32
}{
	{
		input: 1 << 5,
		data: [][]float32{
			randFloats(-1, 1, 1<<5),
			randFloats(-1, 1, 1<<5),
		},
	},
	{
		input: 1 << 7,
		data: [][]float32{
			randFloats(-1, 1, 1<<7),
			randFloats(-1, 1, 1<<7),
		},
	},
	{
		input: 1 << 9,
		data: [][]float32{
			randFloats(-1, 1, 1<<9),
			randFloats(-1, 1, 1<<9),
		},
	},
	{
		input: 1 << 10,
		data: [][]float32{
			randFloats(-1, 1, 1<<10),
			randFloats(-1, 1, 1<<10),
		},
	},
	{
		input: 1 << 11,
		data: [][]float32{
			randFloats(-1, 1, 1<<11),
			randFloats(-1, 1, 1<<11),
		},
	},
	{
		input: 1 << 12,
		data: [][]float32{
			randFloats(-1, 1, 1<<12),
			randFloats(-1, 1, 1<<12),
		},
	},
	{
		input: 1 << 13,
		data: [][]float32{
			randFloats(-1, 1, 1<<13),
			randFloats(-1, 1, 1<<13),
		},
	},
	{
		input: 1 << 14,
		data: [][]float32{
			randFloats(-1, 1, 1<<14),
			randFloats(-1, 1, 1<<14),
		},
	},
	{
		input: 1 << 16,
		data: [][]float32{
			randFloats(-1, 1, 1<<16),
			randFloats(-1, 1, 1<<16),
		},
	},
}

func TestBlas(t *testing.T) {
	for _, v := range table {
		v1 := v.data[0]
		v2 := v.data[1]
		bv1 := blas32.Vector{N: len(v1), Inc: 1, Data: v1}
		bv2 := blas32.Vector{N: len(v2), Inc: 1, Data: v2}
		t.Run(fmt.Sprintf("input_size_%d", v.input), func(t *testing.T) {
			res, err := cosineSimBlas(bv1, bv2)
			assert.NoError(t, err)
			want, err := cosineSimNative(v1, v2)
			assert.NoError(t, err)
			delta := roundFloat(float32(math.Abs(float64(roundFloat(want, precision)-roundFloat(res, precision)))), precision)
			assert.LessOrEqual(t, delta, exp)
		})
	}
}

func TestAsm(t *testing.T) {
	for _, v := range table {
		v1 := v.data[0]
		v2 := v.data[1]
		t.Run(fmt.Sprintf("input_size_%d", v.input), func(t *testing.T) {
			res, err := cosineSimAsm(v1, v2)
			assert.NoError(t, err)
			want, err := cosineSimNative(v1, v2)
			assert.NoError(t, err)
			delta := roundFloat(float32(math.Abs(float64(roundFloat(want, precision)-roundFloat(res, precision)))), precision)
			assert.LessOrEqual(t, delta, exp)
		})
	}
}

func BenchmarkNative(b *testing.B) {
	for _, v := range table {
		v1 := v.data[0]
		v2 := v.data[1]
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cosineSimNative(v1, v2)
			}
		})
	}
}

func BenchmarkBlas(b *testing.B) {
	for _, v := range table {
		v1 := v.data[0]
		v2 := v.data[1]
		bv1 := blas32.Vector{N: len(v1), Inc: 1, Data: v1}
		bv2 := blas32.Vector{N: len(v2), Inc: 1, Data: v2}
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cosineSimBlas(bv1, bv2)
			}
		})
	}
}

func BenchmarkAsm1048(b *testing.B) {
	for _, v := range table {
		v1 := v.data[0]
		v2 := v.data[1]
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cosineSimAsm(v1, v2)
			}
		})
	}
}
