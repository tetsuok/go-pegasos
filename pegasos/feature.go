// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
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

func (f Node) Equal(other Node) bool {
	if f.id != other.id || !close(f.v, other.v) {
		return false
	}
	return true
}

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
	if i >= fv.Len() {
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
	if fv.ptr >= fv.Cap() { // reallocate
		// Allocate double what's needed
		newSlice := make([]Node, fv.ptr*2)

		copy(newSlice, fv.vec)
		fv.vec = newSlice
	}
	fv.vec[fv.ptr] = f
	fv.ptr++
}

func (fv *FeatureVector) Equal(other *FeatureVector) bool {
	if n, n2 := fv.Size(), other.Size(); n != n2 {
		return false
	}
	N := fv.Size()
	for i := 0; i < N; i++ {
		n1 := fv.Index(i)
		n2 := other.Index(i)
		if !n1.Equal(*n2) {
			return false
		}
	}
	return true
}

// Append
// func (fv *FeatureVector) Append(slice []Feature) {
// TODO: implement append slice of Feature to fv.vec.
// }

// TODO: this is too slow. profiling is needed.
func InnerProduct(w []float64, fv *FeatureVector) float64 {
	res := 0.0
	N := fv.Size()
	l := len(w)
	for i := 0; i < N; i++ {
		f := fv.Index(i)
		if f.id >= l {
			break
		}
		res += w[f.id] * f.v
	}
	return res
}
