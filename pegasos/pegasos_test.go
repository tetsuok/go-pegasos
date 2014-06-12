// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
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

var classifierTests = []struct {
	param Param

	// classifier spe
	eta float64
	w   Weights
}{
	{Param{0.1, 2, 1, 1234, "model"}, 0.1, []float64{0.1, 0.2}},
	{Param{0.01, 10, 1, 1234, "model"}, 0.01, []float64{0.1, -0.2}},
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
		if f := hingeLoss(test.w, test.y); !close(test.out, f) {
			t.Errorf("HingeLoss(2.0, 1) = %g, want %g", f, test.out)
		}
	}
}

func TestEncodeClassifier(t *testing.T) {
	for _, test := range classifierTests {
		c := NewClassifier(test.param, []Example{}, 2)
		c.eta = test.eta
		c.w.Swap(test.w)
		bytes := c.Encode()

		var c2 Classifier
		c2.Decode(bytes)
		c2.param.ModelFile = c.param.ModelFile // weird

		if !c.param.Equal(&c2.param) || !close(c.eta, c2.eta) {
			t.Errorf("c2 = %v, expected %v", c2.param, c.param)
		}
		if len(c.w) != len(c2.w) {
			t.Errorf("c2 = %v, expected %v", c2.param, c.param)
		}

		for i := 0; i < len(c.w); i++ {
			if !close(c.w[i], c2.w[i]) {
				t.Errorf("c2 = %v, expected %v", c2.param, c.param)
			}
		}
	}
}
