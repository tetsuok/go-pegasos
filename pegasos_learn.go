// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Usage
//   $ ./pegasos_learn train_file [-m model_file]
//
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tetsuok/go-pegasos/pegasos"
)

var lambda = flag.Float64("lambda", 0.01, "Regularization parameter")
var numIter = flag.Int("t", int(100.0 / *lambda), "Number of iterations")
var blockSize = flag.Int("k", 1, "Size of block for SGD")
var modelFile = flag.String("m", "model", "Model file")

//
// TODO: consider modifying the user interface such as
//
// ./pegasos learn file model
//
// ./pegasos test file model
//
// where "learn" and "test" represent the actions.
//

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Printf("Usage:\npegasos_learn train_file [-m model_file]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	param := pegasos.Param{*lambda, *numIter, *blockSize, *modelFile}
	pegasos.Learn(flag.Arg(0), param)
}
