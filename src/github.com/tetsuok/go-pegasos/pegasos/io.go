// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// reaeder
//
// implements file reader to read files in libsvm format.

func readLines(r io.Reader) {
	rd := bufio.NewReader(r)
	lineNum := 1
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", line)
		lineNum++
	}
}

func ReadTrainingData(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	readLines(file)
}
