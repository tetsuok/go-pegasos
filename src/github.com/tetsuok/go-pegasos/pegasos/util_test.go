// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"testing"
	"math/rand"
)

func TestSelectNext(t *testing.T) {
	rand.Seed(1550503961)
	n := 1000
	for i := 0; i < 10000; i++ {
		r := SelectNext(n)

		if r >= n {
			t.Errorf("r = %d (< %d)", r, n)
		}
	}
}
