// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"fmt"
	"os"
	"strconv"
)

// Tokenize parses a string in the libsvm format.
func Tokenize(line string, lineNum int) (x Example, maxId int) {
	ptr := 0
	begin := ptr
	for {
		if line[ptr] == ' ' {
			break
		}
		ptr++
	}
	y, err := strconv.Atoi(line[begin:ptr])
	if err != nil {
		fmt.Fprintf(os.Stderr, "# y: Illegal line at %d\n", lineNum)
		return Example{}, 0
	}

	ptr++

	fv := NewFeatureVector(50)
	l := len(line)
	for {
		if ptr >= l {
			break
		}
		begin = ptr
		if line[ptr] == '\n' {
			break
		}
		for {
			if line[ptr] == ':' {
				break
			}
			ptr++
		}
		id, err := strconv.Atoi(line[begin:ptr])
		if err != nil {
			fmt.Fprintf(os.Stderr, "# Illegal line at %d\n", lineNum)
			return Example{}, 0
		}
		ptr++
		begin = ptr
		for {
			if ptr == l {
				break
			}
			if line[ptr] == ' ' {
				break
			}
			ptr++
		}
		val, err := strconv.ParseFloat(line[begin:ptr], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "# Illegal line at %d\n", lineNum)
			return Example{}, 0
		}
		fv.PushBack(Node{id, val})
		ptr++

		if id > maxId {
			maxId = id
		}
	}

	return Example{fv, y}, maxId
}
