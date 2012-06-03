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

func readLines(r io.Reader) ([]Example, int) {
	rd := bufio.NewReader(r)
	lineNum := 1
	var data []Example
	maxId := 0
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		x, id, err := Tokenize(strings.TrimRight(line, "\n"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Illegal line at %d\n", lineNum)
			log.Fatal(err)
		}
		data = append(data, x)

		if id > maxId {
			maxId = id
		}

		lineNum++
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

// TODO: Implement reading binary model files.
func OpenModel(model string) (param Param, w Weights, eta float64) {
	f, err := os.Open(param.ModelFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	param = Param{}
	w = NewWeights(10)
	eta = 1.0

	return param, w, eta
}

// TODO: Implement saving the trained model to a file.
func WriteModel(param Param, w Weights, eta float64) {
	f, err := os.Create(param.ModelFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("lambda =", param.Lambda)
	fmt.Println("block size =", param.BlockSize)
	fmt.Println("w =", w)
	fmt.Println("eta =", eta)
}
