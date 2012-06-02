// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"testing"
)

type LossTest struct {
	w   float64
	y   int
	out float64
}

var lossTests = []LossTest{
	{2.0, 1, 0.0},
	{2.0, -1, 3.0},
}

func setupExample(label int) *Example {
	fv := NewFeatureVector(2)
	fv.PushBack(Node{1, 1.0})
	fv.PushBack(Node{3, 2.0})
	return &Example{fv, label}
}

func TestEqualExample(t *testing.T) {
	e := setupExample(1)
	e2 := setupExample(1)
	if !e.Equal(*e2) {
		t.Errorf("%v, want %v", e, e2)
	}
}

func TestObjective(t *testing.T) {
	for _, test := range lossTests {
		if f := HingeLoss(test.w, test.y); !close(test.out, f) {
			t.Errorf("HingeLoss(2.0, 1) = %g, want %g", f, test.out)
		}
	}
}
