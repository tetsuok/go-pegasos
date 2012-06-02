// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"fmt"
)

// binary classification
func HingeLoss(w float64, y int) float64 {
	loss := 1.0 - w*float64(y)
	if loss > 0.0 {
		return loss
	}
	return 0.0
}

// APIs

func Learn(train_file string, param Param) {
	fmt.Println("train_file = ", train_file)
	fmt.Println(param)
}
