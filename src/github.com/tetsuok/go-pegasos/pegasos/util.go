// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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