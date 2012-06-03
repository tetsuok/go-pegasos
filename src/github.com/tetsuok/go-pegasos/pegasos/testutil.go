// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in http://code.google.com/p/go/source/browse/LICENSE.

// This code comes from http://code.google.com/p/go/source/browse/src/pkg/math/all_test.go

package pegasos

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
