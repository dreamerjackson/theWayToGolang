/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
)

func doWork(id int,
	c chan int,done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(
	id int) worker {
	w := worker{
		in: make(chan int),
		done:make(chan bool),
	}
	go doWork(id, w.in,w.done)
	return w
}
//死锁的解决方法1:
func chanDemo() {


	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _,worker := range workers{
		<-worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _,worker := range workers{
		<-worker.done
	}
}

func main() {
	chanDemo()
}

