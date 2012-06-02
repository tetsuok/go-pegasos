# Copyright 2012 Tetsuo Kiso. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

default: all

all: bench

# Do all benchmarks
bench:
	go test -test.v -test.parallel 4 -test.bench ".*"

check:
	go test -test.v -test.parallel 4
