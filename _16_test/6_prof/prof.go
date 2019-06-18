/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	fib "theWayToGolang/_16_test/5_benchmarkTest/1_fibonacci"
)

/*
Use pprof.StartCPUProfile(), pprof.StopCPUProfile() and pprof.WriteHeapProfile(). See the pprof package documentation for more information.

A simple example :

go build
./showfib 40 2>cpu.out
go tool pprof -http=localhost:8080 cpu.out
*/
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : showfib n")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// In real life, don't write to os.Stderr ;)
	// CPU profiling
	pprof.StartCPUProfile(os.Stderr)
	defer pprof.StopCPUProfile()

	fmt.Println(fib.Suite(n))
}

