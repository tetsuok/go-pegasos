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
    $ go build -o pegasos_learn .

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

Under construction.

### Testing with trained model ###

Under construction.

### Reference ####

[1] Shalev-Shwartz, Shai and Singer, Yoram and Srebro,
Nathan. Pegasos: Primal Estimated sub-GrAdient SOlver for SVM.
In Proceedings of the 24th international conference on Machine learning
(ICML). 2007. pages 807-814.
