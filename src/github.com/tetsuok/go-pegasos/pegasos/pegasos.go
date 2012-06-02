// Copyright 2012 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// binary classification
func hingeLoss(w float64, y int) float64 {
	loss := 1.0 - w*float64(y)
	if loss > 0.0 {
		return loss
	}
	return 0.0
}

func HingeLoss(w Weights, fv *FeatureVector, y int) float64 {
	score := InnerProduct(w, fv)
	return hingeLoss(score, y)
}

// Example represents each example (x, y) in training data where x is
// a feature vector representation and y is a label: -1 or +1.
type Example struct {
	fv    *FeatureVector
	label int
}

type MissedExample struct {
	id int												// an index of missed example.
	w  float64										// score to calculate eta_k /k \sum y x
}

func (e Example) Equal(other Example) bool {
	if e.label != other.label {
		return false
	}
	if !e.fv.Equal(other.fv) {
		return false
	}
	return true
}

type Classifier struct {
	param Param
	examples []Example
	w Weights
	eta float64
}

func NewClassifier(param Param, examples []Example, dim int) *Classifier {
	return &Classifier{param, examples, NewWeights(dim+1), 0.0}
}

func (c *Classifier) SetEta(t int) {
	c.eta = 1.0 / (c.param.Lambda * float64(t + 2))
}

func (c *Classifier) Eta() float64 { return c.eta }

// func (c *Classifier)CalcMissedExamples() {

// }

func (c *Classifier) Project() {
	norm := c.w.L2Norm()
	l := math.Sqrt(c.param.Lambda)
	coeff := 1.0 / (l * norm)
	if 1.0 > coeff {
		c.w.Scale(coeff)
	}
}

func (c *Classifier) CalcObjective(missed []MissedExample) float64 {
	norm := c.w.SquareNorm() * 0.5 * c.param.Lambda

	loss := 0.0
	for _, e := range missed {
		loss += HingeLoss(c.w, c.examples[e.id].fv, c.examples[e.id].label)
	}
	loss /= float64(c.param.BlockSize)
	return norm + loss
}

// Open model
func (c *Classifier) Open() {
}

// Save model to a file
func (c *Classifier) Save() {
	SaveModel(c.param, c.w, c.eta)
}

// APIs

func Learn(train_file string, param Param) {
	// fmt.Println("train_file = ", train_file)
	fmt.Println(param)
	ReadTrainingData(train_file)
}
