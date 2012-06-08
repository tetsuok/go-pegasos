// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Usage
//   $ ./pegasos_learn -m model_file train_file
//
// on the fly testing
//   $ ./pegasos_learn -test test_file train_file
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
var modelFile = flag.String("m", "", "Model file")
var testFile = flag.String("test", "", "Test file. This option enables on the fly testing")
var seed = flag.Int64("r", 1550503961, "Random seed")

func usage() {
	fmt.Printf("Usage:\npegasos_learn -m model_file train_file\n")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
		os.Exit(1)
	}
	trainFile := flag.Arg(0)
	param := pegasos.Param{*lambda, *numIter, *blockSize, *seed, *modelFile}

	hasModel := len(*modelFile) > 0
	hasTest := len(*testFile) > 0
	switch {
	case hasModel && !hasTest:
		pegasos.Learn(trainFile, param)
	case !hasModel && hasTest:
		// on the fly training and testing.
		pegasos.LearnAndClassify(trainFile, *testFile, param)
	case hasModel && hasTest:
		// on the fly training and testing but we save trained model.
		pegasos.LearnAndClassify(trainFile, *testFile, param)
	default:
		usage()
		os.Exit(1)
	}
}
