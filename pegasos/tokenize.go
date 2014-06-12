// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// An example of tokenize strings

package pegasos

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrSyntax = errors.New("invalid syntax")

// NumError same as strconv.atoi.go
type NumError struct {
	Func string
	Num  string
	Err  error
}

func (e *NumError) Error() string {
	return "tokenize." + e.Func + ": " + e.Num + ": " + e.Err.Error()
}

func syntaxError(fn, str string) *NumError {
	return &NumError{fn, str, ErrSyntax}
}

func tokenizeNode(s string) (v *Node, err error) {
	const fnTokenizeNode = "tokenizeNode"
	if len(s) == 0 {
		return nil, syntaxError(fnTokenizeNode, "empty string")
	}

	sep := uint8(':')
	id := 0
	val := 0.0

	for i := 0; i < len(s); i++ {
		if s[i] != sep {
			continue
		}
		id, err = strconv.Atoi(s[0:i])
		if err != nil {
			goto Error
		}
		val, err = strconv.ParseFloat(s[i+1:], 64)
		if err != nil {
			goto Error
		}
		return &Node{id, val}, nil
	}

Error:
	return nil, syntaxError(fnTokenizeNode, s)
}

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

// func DebugStringList(l *list.List) string {
// 	str := "["
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		if e != l.Front() {
// 			str += " "
// 		}
// 		if e.Value != nil {
// 			str += fmt.Sprint(e.Value)
// 		}
// 	}
// 	return str + "]"
// }
