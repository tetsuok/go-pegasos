// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// an implementation of weight vector.

package pegasos

import "math"

type Weights []float64

func NewWeights(dim int) Weights {
	return Weights(make([]float64, dim))
}

func (w *Weights) Get(i int) float64 { return (*w)[i] }
func (w *Weights) Len() int          { return len(*w) }

func (w *Weights) Scale(f float64) {
	for i := 0; i < len(*w); i++ {
		(*w)[i] = (*w)[i] * f
	}
}

func (w *Weights) Add(fv *FeatureVector, c float64) {
	N := fv.Size()
	for i := 0; i < N; i++ {
		f := fv.Index(i)
		(*w)[f.id] += f.v * c
	}
}

func (w *Weights) SquareNorm() float64 {
	res := 0.0
	for i := 0; i < len(*w); i++ {
		res += (*w)[i] * (*w)[i]
	}
	return res
}

func (w *Weights) L2Norm() float64 {
	return math.Sqrt(w.SquareNorm())
}

func (w *Weights) Swap(a []float64) {
	*w = Weights(a)
}
