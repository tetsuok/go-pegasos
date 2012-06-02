// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"testing"
)

type WeightsTest struct {
	w Weights
}

// TODO: refactor test cases.

func TestLen(t *testing.T) {
	w := &Weights{2.0, 1.0}
	if l := w.Len(); l != 2 {
		t.Errorf("w.Len() = %d, want %d", l, 2)
	}
}

func TestScale(t *testing.T) {
	w := &Weights{2.0, 3.0}
	w.Scale(2.0)

	expected := &Weights{4.0, 6.0}
	for i := 0; i < w.Len(); i++ {
		if !close(w.Get(i), expected.Get(i)) {
			t.Errorf("w[%d] = %g (expected[%d] = %g)", i, w.Get(i), i, expected.Get(i))
		}
	}
}

func TestSquareNorm(t *testing.T) {
	w := &Weights{2.0, 3.0}

	if norm := w.SquareNorm(); !close(norm, 13.0) {
		t.Errorf("w.SquareNorm() = %g, want %g", norm, 13.0)
	}
}

func TestL2Norm(t *testing.T) {
	w := &Weights{4.0, 3.0}

	if norm := w.L2Norm(); !close(norm, 5.0) {
		t.Errorf("w.L2Norm() = %g, want %g", norm, 5.0)
	}
}

func TestSwap(t *testing.T) {
	w := &Weights{4.0, 3.0}

	a := []float64{3.0, 1.0}
	w.Swap(a)
	for i := 0; i < w.Len(); i++ {
		if !close(w.Get(i), a[i]) {
			t.Errorf("w[%d] = %g (expected[%d] = %g)", i, w.Get(i), i, a[i])
		}
	}
}

// Benchmarking stuff

func BenchmarkScale(b *testing.B) {
	w := &Weights{4.0, 3.0, 2.0, 2.0, 1.0, 4.0}
	for i := 0; i < b.N; i++ {
		w.Scale(100.0)
	}
}

func BenchmarkSquareNorm(b *testing.B) {
	w := &Weights{4.0, 3.0, 2.0, 2.0, 1.0, 4.0}
	for i := 0; i < b.N; i++ {
		w.SquareNorm()
	}
}

func BenchmarkL2Norm(b *testing.B) {
	w := &Weights{4.0, 3.0, 2.0, 2.0, 1.0, 4.0}
	for i := 0; i < b.N; i++ {
		w.L2Norm()
	}
}
