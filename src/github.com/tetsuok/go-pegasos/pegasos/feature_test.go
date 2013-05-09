// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"math/rand"
	"sort"
	"testing"
)

func setupFeatureVector() *FeatureVector {
	fv := NewFeatureVector(2)
	fv.PushBack(Node{1, 1.0})
	fv.PushBack(Node{3, 2.0})
	return fv
}

func TestNode(t *testing.T) {
	f := Node{1, 4.0}
	if f.id != 1 {
		t.Errorf("f.id = %d, want %d", f.id, 1)
	}
	if !close(f.v, 4.0) {
		t.Errorf("f.v = %g, want %g", f.v, 4.0)
	}
}

func TestNodeEqual(t *testing.T) {
	n := Node{1, 3.0}
	n2 := Node{1, 3.0}
	if !n.Equal(n2) {
		t.Errorf("n = %v, want %v", n, n2)
	}
}

func TestFeatureVectorLen(t *testing.T) {
	l := 10
	fv := NewFeatureVector(l)
	if fv.Len() != l {
		t.Errorf("fv.Len() = %d, want %d", fv.Len(), l)
	}
}

func TestFeatureVectorSize(t *testing.T) {
	fv := NewFeatureVector(2)
	fv.PushBack(Node{1, 4.0})
	if fv.Size() != 1 {
		t.Errorf("fv.Size() = %d, want %d", fv.Size(), 1)
	}
}

func TestFeatureIndex(t *testing.T) {
	fv := NewFeatureVector(2)
	f := Node{1, 4.0}
	fv.PushBack(f)
	if fi := fv.Index(0); *fi != f {
		t.Errorf("fi = %s, want %s", fi, f)
	}
}

func TestFeatureVectorEqual(t *testing.T) {
	fv1 := setupFeatureVector()
	fv2 := setupFeatureVector()

	if !fv1.Equal(fv2) {
		t.Errorf("%v, want %v", fv1, fv2)
	}
}

func TestFeatureInnerProduct(t *testing.T) {
	fv := NewFeatureVector(2)
	fv.PushBack(Node{0, 1.0})
	fv.PushBack(Node{1, 4.0})
	w := []float64{2.0, 4.0}
	if v := InnerProduct(w, fv); !close(v, 18.0) {
		t.Errorf("InnerProduct(w, fv) = %g, want %g", v, 18.0)
	}
}

func TestInnerProductBoundaryCheck(t *testing.T) {
	fv := NewFeatureVector(2)
	fv.PushBack(Node{0, 1.0})
	fv.PushBack(Node{1, 4.0})
	fv.PushBack(Node{2, 4.0})
	w := []float64{2.0, 4.0}
	if v := InnerProduct(w, fv); !close(v, 18.0) {
		t.Errorf("InnerProduct(w, fv) = %g, want %g", v, 18.0)
	}
}

func makeSparseFeatureVector(size int) *FeatureVector {
	fv := NewFeatureVector(size)

	// Sample sparse feature ids.
	ids := make([]int, size/50)
	for i := 0; i < len(ids); i++ {
		ids[i] = rand.Intn(size)
	}
	sort.Ints(ids)

	for i := 0; i < len(ids); i++ {
		v := rand.Float64()
		fv.PushBack(Node{ids[i], v})
	}
	return fv
}

func BenchmarkInnerProduct(b *testing.B) {
	b.StopTimer()
	rand.Seed(10)

	l := 10000
	fv := makeSparseFeatureVector(l)

	w := make([]float64, l)
	for i := 0; i < l; i++ {
		w[i] = rand.Float64()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		InnerProduct(w, fv)
	}
}
