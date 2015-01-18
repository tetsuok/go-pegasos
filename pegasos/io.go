// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"bufio"
	"io"
	"log"
	"os"
)

func readLines(r io.Reader) ([]Example, int) {
	scanner := bufio.NewScanner(r)
	lineNum := 1
	data := make([]Example, 0, 1000)
	maxId := 0
	for scanner.Scan() {
		line := scanner.Text()
		x, id := tokenize(line, lineNum)
		data = append(data, x)

		if id > maxId {
			maxId = id
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data, maxId
}

func ReadTrainingData(filename string) ([]Example, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return readLines(file)
}
