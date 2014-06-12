// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Param represents hyperparameters used in Pegasos.
type Param struct {
	// hyperparameter used in Pegasos.
	Lambda float64

	// The number of iterations
	NumIter int

	// Pegasos's k (size of A_t)
	BlockSize int

	// seed to perform random sampling from the training data.
	Seed int64

	// Model file to be used in loading and saving to a file.
	ModelFile string
}

// Convert the internal parameters into bytes.
func (p *Param) Encode() []byte {
	buf := p.Buffer()
	return buf.Bytes()
}

func (p *Param) Buffer() *bytes.Buffer {
	buf := new(bytes.Buffer)
	encode(buf, p.Lambda)

	var it int32 = int32(p.NumIter)
	var k int32 = int32(p.BlockSize)
	encode(buf, it)
	encode(buf, k)
	return buf
}

// Decode decodes bytes internal parameters again.
func (p *Param) Decode(data []byte) *bytes.Buffer {
	buf := bytes.NewBuffer(data)
	decode(buf, &p.Lambda)

	var it int32
	var k int32
	decode(buf, &it)
	decode(buf, &k)
	p.NumIter = int(it)
	p.BlockSize = int(k)
	return buf
}

func encode(b *bytes.Buffer, data interface{}) {
	if err := binary.Write(b, binary.LittleEndian, data); err != nil {
		log.Fatal(err)
	}
}

func decode(b *bytes.Buffer, data interface{}) {
	if err := binary.Read(b, binary.LittleEndian, data); err != nil {
		log.Fatal(err)
	}
}

func (p *Param) Equal(other *Param) bool {
	if !close(p.Lambda, other.Lambda) || p.NumIter != other.NumIter || p.BlockSize != other.BlockSize || p.ModelFile != other.ModelFile {
		return false
	}
	return true
}
