go-pegasos
==========

An implementation of the Pegasos algorithm [1] for solving Support Vector Machines in Go.

Build Instructions
------------------

### Software Requirements ###

* [Go](http://golang.org/)

### Get the code ###

    $ git clone git://github.com/tetsuok/go-pegasos.git

### Compilation ###

    $ export GOPATH=$(pwd)
    $ go build github.com/tetsuok/go-pegasos/pegasos_learn
    $ go build github.com/tetsuok/go-pegasos/pegasos_classify

or simply run

    $ ./build.sh

`build.sh` is a wrapper of these commands.


### Testing ###

    $ export GOPATH=$(pwd)
    $ go test github.com/tetsuok/go-pegasos/pegasos

If you want to run testing including benchmarks, use `check.sh`

    $ ./check.sh


Usage
-----

### Data format ###

go-pegasos accepts the same representation of training data as
[SVMlight](http://svmlight.joachims.org/) uses. This format has
potential to handle large sparse feature vectors.

### Training ###

    $ ./pegasos_learn -m model_file train_file

Please note "-m" is required to save the trained model.

#### Options #####

* -k INT: number of block size.
* -lambda FLOAT: Regularization parameter
* -m STRING: model file
* -r INT: seed
* -t INT: number of iterations
* -test STRING: If you set a test file, you can do training and testing at a time.

### Testing with trained model ###

    $ ./pegasos_test test_file model_file

### Reference ####

[1] Shalev-Shwartz, Shai and Singer, Yoram and Srebro,
Nathan. Pegasos: Primal Estimated sub-GrAdient SOlver for SVM.
In Proceedings of the 24th international conference on Machine learning
(ICML). 2007. pages 807-814.
