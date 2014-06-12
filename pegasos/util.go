// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"math/rand"
)

func SelectNext(n int) int {
	return rand.Intn(n)
}
