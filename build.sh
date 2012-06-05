#!/bin/sh
# Copyright 2012 Tetsuo Kiso. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -x
OPT=-x
export GOPATH=$(pwd)
go build $OPT github.com/tetsuok/go-pegasos/pegasos_learn
go build $OPT github.com/tetsuok/go-pegasos/pegasos_classify
