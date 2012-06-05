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
	id int     // an index of missed example.
	w  float64 // score to calculate eta_k /k \sum y x
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
	param    Param
	examples []Example
	eta      float64
	w        Weights
}

func NewClassifier(param Param, examples []Example, dim int) *Classifier {
	return &Classifier{param, examples, 0.0, NewWeights(dim + 1)}
}

func (c *Classifier) SetEta(t int) {
	c.eta = 1.0 / (c.param.Lambda * float64(t+2))
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
func (c *Classifier) Open(model string) {
	c.param, c.w, c.eta = OpenModel(model)
}

// Save model to a file
func (c *Classifier) Save() {
	WriteModel(c.param, c.w, c.eta)
}

// APIs

func Learn(trainFile string, param Param) {
	rand.Seed(1234)
	start := time.Now()

	fmt.Printf("Reading %s ... ", trainFile)
	examples, dim := ReadTrainingData(trainFile)

	fmt.Printf("Done!. Elapsed time %s\n", time.Since(start))

	classifier := NewClassifier(param, examples, dim)

	numExamples := len(examples)
	fmt.Println("Dimension =", classifier.w.Len())
	fmt.Println("# of training data =", numExamples)

	for t := 0; t < param.NumIter; t++ {
		fmt.Println("Iteration =", t)
		// eta := 1.0 / (param.Lambda * float64(t + 2))
		classifier.SetEta(t)

		// TODO: use make and resize properly.
		var missedExamples []MissedExample

		// Set up A_t
		for k := 0; k < param.BlockSize; k++ {
			r := SelectNext(numExamples)

			// Compute A_t^+
			loss := HingeLoss(classifier.w, examples[r].fv, examples[r].label)
			if loss > 0.0 {
				// TODO: This is too slow; replace Append() described in "Effective Go".
				missedExamples = append(missedExamples,
					MissedExample{r, classifier.eta * float64(examples[r].label) / float64(param.BlockSize)})
			}
		}

		// Subgradient
		classifier.w.Scale(1.0 - classifier.eta*param.Lambda)
		for _, missed := range missedExamples {
			classifier.w.Add(examples[missed.id].fv, missed.w)
		}

		classifier.Project()

		// Compute objective
		fmt.Println("objective =", classifier.CalcObjective(missedExamples))
	}

	classifier.Save()

	fmt.Printf("Done!. Elapsed time %s\n", time.Since(start))
}
