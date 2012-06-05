// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Usage
//   $ ./classify test_file model_file

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tetsuok/go-pegasos/pegasos"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Printf("Usage:\n./classify test_file model_file\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	pegasos.Classify(flag.Arg(0), flag.Arg(1))
}
