default: all

all: bench

# Do all benchmarks
bench:
	go test -test.v -test.parallel 4 -test.bench ".*"

check:
	go test -test.v -test.parallel 4
