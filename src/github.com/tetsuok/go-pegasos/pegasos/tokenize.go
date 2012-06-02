// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// An example of tokenize strings

package pegasos

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func Tokenize(s string) (x Example, err error) {
	fv := NewFeatureVector(3)
	seq := strings.Split(s, " ")
	y, err := strconv.Atoi(seq[0])

	if err != nil || (y != -1 && y != 1) {
		fmt.Fprintf(os.Stderr, "Invalid label = %d\n", y)
		return Example{}, err
	}

	for _, a := range seq[1:] {
		node, err := tokenizeNode(a)
		if err != nil {
			return Example{}, err
		}
		fv.PushBack(*node)
	}
	return Example{fv, y}, nil
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
