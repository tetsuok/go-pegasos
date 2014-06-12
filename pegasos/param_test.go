// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"testing"
)

var paramTests = []struct {
	lambda    float64
	iter      int
	blockSize int
}{
	{0.1, 2, 1},
	{0.01, 10, 1},
}

func TestEncodeParam(t *testing.T) {
	for _, test := range paramTests {
		param := Param{test.lambda, test.iter, test.blockSize, 0, ""}
		bytes := param.Encode()
		var param2 Param
		param2.Decode(bytes)
		if !param.Equal(&param2) {
			t.Errorf("param2 = %v, expected %v", param2, param)
		}
	}
}
