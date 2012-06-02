// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"fmt"
	"math/rand"
)

// binary classification
func HingeLoss(w float64, y int) float64 {
	loss := 1.0 - w*float64(y)
	if loss > 0.0 {
		return loss
	}
	return 0.0
}

// Example represents each example (x, y) in training data where x is
// a feature vector representation and y is a label: -1 or +1.
type Example struct {
	fv    *FeatureVector
	label int
}

func (e Example) Equal(other Example) bool {
	if e.label != other.label {
		return false
	}
	if !e.fv.Equal(other.fv) {
		return false
	}
	return true
}

// APIs

func Learn(train_file string, param Param) {
	// fmt.Println("train_file = ", train_file)
	fmt.Println(param)
	ReadTrainingData(train_file)
}
