#!/bin/sh
# Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -x

# Just testing.
# go test -test.v -test.parallel 4 github.com/tetsuok/go-pegasos/pegasos

# Testing + benchmarking
go test -test.v -test.parallel 4 -test.bench ".*" github.com/tetsuok/go-pegasos/pegasos
