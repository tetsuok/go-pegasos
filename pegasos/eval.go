// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import "fmt"

type Eval struct {
	// confusion matrix
	truePositive  int
	trueNegative  int
	falsePositive int
	falseNegative int
	numInstance   int
}

func (e *Eval) Evaluate(ans, predict int) {
	switch {
	case ans == 1 && predict == 1:
		e.truePositive++
	case ans == -1 && predict == -1:
		e.trueNegative++
	case ans == -1 && predict == 1:
		e.falsePositive++
	case ans == 1 && predict == -1:
		e.falseNegative++
	}
	e.numInstance++
}

// in percentage

func (e *Eval) Accuracy() float64 {
	return float64(e.numInstance - (e.falsePositive + e.falseNegative)) * 100.0 / float64(e.numInstance)
}

func (e *Eval) Precision() float64 {
	return float64(e.truePositive) / float64(e.truePositive + e.falsePositive)
}

func (e *Eval) Precision100() float64 {
	return e.Precision() * 100.0
}

func (e *Eval) Recall() float64 {
	return float64(e.truePositive) / float64(e.truePositive + e.falseNegative)
}

func (e *Eval) Recall100() float64 {
	return e.Recall() * 100.0
}

func (e *Eval) F1() float64 {
	p := e.Precision()
	r := e.Recall()
	return (2 * p * r) / (p + r)
}

func (e *Eval) F1100() float64 {
	return e.F1() * 100.0
}

func (e *Eval) String() string {
	res := fmt.Sprintf("Accuracy %g%% ", e.Accuracy())
	res += fmt.Sprintf("(%d/%d) ", e.numInstance - (e.falsePositive + e.falseNegative), e.numInstance)
	res += fmt.Sprintf("tp %d tn %d fp %d fn %d", e.truePositive, e.trueNegative, e.falsePositive, e.falseNegative)
	return res
}
