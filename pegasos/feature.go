// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

// an simple implementation of feature vector in go

type Node struct {
	id int
	v  float64
}

func (f Node) Id() int        { return f.id }
func (f Node) Value() float64 { return f.v }

type FeatureVector struct {
	vec []Node
	ptr int // index of new element to be inserted
}

func (fv *FeatureVector) Len() int  { return len(fv.vec) }
func (fv *FeatureVector) Cap() int  { return cap(fv.vec) }
func (fv *FeatureVector) Size() int { return fv.ptr }

func (fv *FeatureVector) Index(i int) *Node {
	if i < 0 {
		panic("invalid id")
	}
	if i > fv.Len() {
		panic("index exceeds the size of feature vector")
	}
	return &fv.vec[i]
}

// Alloc
func NewFeatureVector(size int) *FeatureVector {
	v := make([]Node, size)
	fv := FeatureVector{vec: v}
	return &fv
}

func (fv *FeatureVector) PushBack(f Node) {
	l := fv.Len()

	if l > fv.Cap() { // reallocate
		// Allocate double what's needed
		newSlice := make([]Node, l*2)

		copy(newSlice, fv.vec)
		fv.vec = newSlice
	}
	fv.vec[fv.ptr] = f
	fv.ptr++
}

// Append
// func (fv *FeatureVector) Append(slice []Feature) {
// TODO: implement append slice of Feature to fv.vec.
// }

// TODO: this is too slow. profiling is needed.
func InnerProduct(w []float64, fv *FeatureVector) float64 {
	res := 0.0
	N := fv.Size()
	for i := 0; i < N; i++ {
		f := fv.Index(i)
		res += w[f.id] * f.v
	}
	return res
}
