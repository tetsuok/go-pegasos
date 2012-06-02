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
	"strings"
)

// reaeder
//
// implements file reader to read files in libsvm format.

func readLines(r io.Reader) []Example {
	rd := bufio.NewReader(r)
	lineNum := 1
	var data []Example
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		x, err := Tokenize(strings.TrimRight(line, "\n"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Illegal line at %d\n", lineNum)
			log.Fatal(err)
		}
		data = append(data, x)

		lineNum++
	}
	return data
}

func ReadTrainingData(filename string) []Example {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return readLines(file)
}
