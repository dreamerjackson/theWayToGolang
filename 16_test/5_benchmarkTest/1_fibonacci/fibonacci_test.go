/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package fibonacci

import (
	"fmt"
	"testing"
)





func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fibonacci(30)
	}
}

func BenchmarkSuite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Suite(30)
	}
}
var result int


//这个问题似乎已经不存在
func BenchmarkFibComplete(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fibonacci(30)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

/*
without any benchmarks : go test .
with benchmarks (time) : go test . -bench .
with benchmarks (time and memory) : go test . -bench . -benchmem
The argument following -bench is a regular expression.
All benchmark functions whose names match are executed.
The . in the previous examples isn't the current directory but a pattern matching all tests.
To run a specific benchmark, use the regexp : -bench Suite (means everything containing Suite).
Useful tip : see ResetTimer() to ignore test setup in measures,
see also StopTimer() and StartTimer(): https://golang.org/pkg/testing/#B.ResetTimer
It's possible to compare benchmarks with an external tool :

go get -u golang.org/x/tools/cmd/benchcmp

go test ./fibonacci -bench . -benchmem > old.txt
(do some changes in the code)
go test ./fibonacci -bench . -benchmem > new.txt

~/go/bin/benchcmp old.txt new.txt

go test . \
  -bench BenchmarkSuite \
  -benchmem \
  -cpuprofile=cpu.out \
  -memprofile=mem.out


proff
Profiling benchmarks
Get profiling data from the benchmarks:

CPU profiling using -cpuprofile=cpu.out
Memory profiling using -benchmem -memprofile=mem.out
An example with both :

go test ./fibonacci \
  -bench BenchmarkSuite \
  -benchmem \
  -cpuprofile=cpu.out \
  -memprofile=mem.out
CPU and memory profiling data from benchmarks are always stored in two separate files and will be analysed separately.

Viewing profiling data
There are two way to exploit profiling data with standard go tools :

through command line : go tool pprof cpu.out
with a browser : go tool pprof -http=localhost:8080 cpu.out
The View menu :

Top : ordered list of functions sorted by their duration/memory consumption.
Graph : function call tree, with time/memory annotations.
Flamegraph : self-explanatory
others...
Graph and flamegraph are rather similar, but there is a major difference :
flamegraph shows sampled call stack with time/memory data (it's a tree),
while graph could have multiple path converging to the same function (it's not a tree).
The time/memory data associated to functions having multiple paths converging to is an aggregation.
Both are useful, just choose the good one for your case.
*/

/*

var knownFibonacci = map[int]int{
	1: 1,
	2: 1,
	3: 2,
	4: 3,
	5: 5,
	6: 8,
}

A note on compiler optimisations
Before concluding I wanted to highlight that to be completely accurate, any benchmark should be careful to avoid compiler optimisations eliminating the function under test and artificially lowering the run time of the benchmark.

var result int

func BenchmarkFibComplete(b *testing.B) {
        var r int
        for n := 0; n < b.N; n++ {
                // always record the result of Fib to prevent
                // the compiler eliminating the function call.
                r = Fib(10)
        }
        // always store the result to a package level variable
        // so the compiler cannot eliminate the Benchmark itself.
        result = r
}

*/



func TestFibonacci(t *testing.T) {
	for n, result := range knownFibonacci {
		description := fmt.Sprintf("Fib(%d)", n)
		t.Run(description, func(t *testing.T) {
			fib := Fibonacci(n)
			if fib != result {
				t.Error(fmt.Sprintf("Expected Fibonnaci(%d) == %d, but was %d", n, result, fib))
			}
		})
	}
}

func TestSuite(t *testing.T) {
	expected := "Fib(1)\t= 1\nFib(2)\t= 1\nFib(3)\t= 2"
	suite := Suite(3)
	if suite != expected {
		t.Errorf("Expected [%s] but was [%s]", expected, suite)
	}
}

var knownFibonacci = map[int]int{
	1: 1,
	2: 1,
	3: 2,
	4: 3,
	5: 5,
	6: 8,
}

