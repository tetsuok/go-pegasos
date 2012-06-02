// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"testing"
)

func tolerance(a, b, e float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}

	if a != 0 {
		e = e * a
		if e < 0 {
			e = -e
		}
	}
	return d < e
}

// func kindaclose(a, b float64) bool { return tolerance(a, b, 1e-8) }
func close(a, b float64) bool { return tolerance(a, b, 1e-14) }

type LossTest struct {
	w   float64
	y   int
	out float64
}

var lossTests = []LossTest{
	{2.0, 1, 0.0},
	{2.0, -1, 3.0},
}

func TestObjective(t *testing.T) {
	for _, test := range lossTests {
		if f := HingeLoss(test.w, test.y); !close(test.out, f) {
			t.Errorf("HingeLoss(2.0, 1) = %g, want %g", f, test.out)
		}
	}
}
