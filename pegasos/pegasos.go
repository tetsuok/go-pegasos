// Copyright 2012-2014 Tetsuo Kiso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pegasos

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
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

func (e Example) Equal(other Example) bool {
	if e.label != other.label {
		return false
	}
	if !e.fv.Equal(other.fv) {
		return false
	}
	return true
}

type MissedExample struct {
	id int     // an index of missed example.
	w  float64 // score to calculate eta_k /k \sum y x
}

type MissedExamples struct {
	examples []MissedExample
	lastId   int // point to last id of slices "examples"
}

func NewMissedExamples(size int) *MissedExamples {
	return &MissedExamples{make([]MissedExample, size), 0}
}

func (m *MissedExamples) SetId(id int) {
	m.examples[m.lastId].id = id
}

func (m *MissedExamples) SetValue(v float64) {
	m.examples[m.lastId].w = v
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

// Note: we won't encode training examples.
func (c *Classifier) Encode() []byte {
	buf := c.param.Buffer()
	c.writeBytes(buf)
	return buf.Bytes()
}

func (c *Classifier) Decode(data []byte) {
	buf := c.param.Decode(data)
	decode(buf, &c.eta)

	// decode a weight vector
	var l uint64
	decode(buf, &l)

	c.w = make([]float64, l)
	for i := 0; i < len(c.w); i++ {
		decode(buf, &c.w[i])
	}
}

func (c *Classifier) writeBytes(buf *bytes.Buffer) {
	encode(buf, c.eta)

	// Encode the weight vector
	encode(buf, uint64(len(c.w)))
	for _, v := range c.w {
		encode(buf, v)
	}
}

func (c *Classifier) SetEta(t int) {
	c.eta = 1.0 / (c.param.Lambda * float64(t+2))
}

func (c *Classifier) Eta() float64 { return c.eta }

// func (c *Classifier)CalcMissedExamples() {

// }

func (c *Classifier) Predict(fv *FeatureVector) int {
	s := InnerProduct(c.w, fv)
	if s > 0.0 {
		return 1
	}
	return -1
}

func (c *Classifier) Project() {
	norm := c.w.L2Norm()
	l := math.Sqrt(c.param.Lambda)
	coeff := 1.0 / (l * norm)
	if 1.0 > coeff {
		c.w.Scale(coeff)
	}
}

func (c *Classifier) CalcObjective(examples []Example) float64 {
	norm := c.w.SquareNorm() * 0.5 * c.param.Lambda

	totalLoss := 0.0
	for _, e := range examples {
		totalLoss += HingeLoss(c.w, e.fv, e.label)
	}
	totalLoss /= float64(len(examples))
	return norm + totalLoss
}

func (c *Classifier) Objective(totalLoss float64) float64 {
	norm := c.w.SquareNorm() * 0.5 * c.param.Lambda
	return norm + totalLoss
}

// Open model
func (c *Classifier) ReadModel(model string) error {
	data, err := ioutil.ReadFile(model)
	if err != nil {
		log.Fatal(err)
	}
	c.Decode(data)
	return err
}

// Save model to a file
func (c *Classifier) WriteModel(model string) error {
	return ioutil.WriteFile(model, c.Encode(), 0644)
}

// APIs

func Learn(trainFile string, param Param) {
	rand.Seed(param.Seed)
	start := time.Now()

	fmt.Printf("Reading %s ... ", trainFile)
	examples, dim := ReadTrainingData(trainFile)

	fmt.Printf("Done!. Elapsed time %s\n", time.Since(start))

	classifier := NewClassifier(param, examples, dim)
	trainClassifier(examples, classifier)

	if len(param.ModelFile) > 0 {
		classifier.WriteModel(param.ModelFile)
	}

	fmt.Printf("Done!. Elapsed time %s\n", time.Since(start))
}

// Classify classifies test examples with trained model.
func Classify(testFile string, model string) {
	start := time.Now()

	var classifier Classifier
	classifier.ReadModel(model)
	fmt.Printf("Model loaded %s\n", time.Since(start))

	classifyTestData(testFile, &classifier)

	// TODO: Compute accuracy and recall.

	fmt.Printf("Done!. Elapsed time %s\n", time.Since(start))
}

func LearnAndClassify(trainFile, testFile string, param Param) {
	rand.Seed(param.Seed)
	start := time.Now()

	fmt.Printf("Reading %s ... ", trainFile)
	examples, dim := ReadTrainingData(trainFile)
	fmt.Printf("Done!. Elapsed time %s\n", time.Since(start))

	classifier := NewClassifier(param, examples, dim)
	trainClassifier(examples, classifier)

	classifyTestData(testFile, classifier)

	if len(param.ModelFile) > 0 {
		classifier.WriteModel(param.ModelFile)
	}

	fmt.Printf("Done!. Elapsed time %s\n", time.Since(start))
}

func trainClassifier(examples []Example, classifier *Classifier) {
	numExamples := len(examples)
	fmt.Println("Number of features =", classifier.w.Len())
	fmt.Println("Number of training examples =", numExamples)

	for t := 0; t < classifier.param.NumIter; t++ {
		classifier.SetEta(t)

		missedExamples := NewMissedExamples(classifier.param.BlockSize)

		// Set up A_t
		for k := 0; k < classifier.param.BlockSize; k++ {
			r := SelectNext(numExamples)

			// Compute A_t^+
			loss := HingeLoss(classifier.w, examples[r].fv, examples[r].label)
			if loss > 0.0 {
				v := classifier.eta * float64(examples[r].label) / float64(classifier.param.BlockSize)
				missedExamples.SetId(r)
				missedExamples.SetValue(v)
				missedExamples.lastId++
			}
		}

		// Subgradient
		classifier.w.Scale(1.0 - classifier.eta*classifier.param.Lambda)

		// Compute w_{t+1/2} = (1 - eta_t * lambda) w_t + (eta_t / k) \sum_{x, y} \in A_t y x
		for i := 0; i < missedExamples.lastId; i++ {
			missed := missedExamples.examples[i]
			classifier.w.Add(examples[missed.id].fv, missed.w)
		}

		classifier.Project()

		// This should be used only for experiments since this gets slow.
		// Compute objective
		// fmt.Fprintf(os.Stderr, "objective = %g\n", classifier.CalcObjective(examples))
	}
}

func classifyTestData(testFile string, classifier *Classifier) {
	reader, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	lineNum := 1
	eval := Eval{}
	totalLoss := 0.0

	for scanner.Scan() {
		line := scanner.Text()

		x, _ := Tokenize(line, lineNum)

		if err != nil {
			fmt.Fprintf(os.Stderr, "# Illegal line at %d\n", lineNum)
			continue
		}

		y := classifier.Predict(x.fv)
		eval.Evaluate(x.label, y)
		totalLoss += HingeLoss(classifier.w, x.fv, x.label)
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	totalLoss /= float64(lineNum)

	fmt.Println(eval.String())
	fmt.Fprintf(os.Stderr, "objective = %g\n", classifier.Objective(totalLoss))
}
