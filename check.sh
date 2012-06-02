#!/bin/sh
# Copyright 2012 Tetsuo Kiso. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -x
export GOPATH=$(pwd)

# Just testing.
# go test -test.v -test.parallel 4 github.com/tetsuok/go-pegasos/pegasos

# Testing + benchmarking
go test -test.v -test.parallel 4 -test.bench ".*" github.com/tetsuok/go-pegasos/pegasos
