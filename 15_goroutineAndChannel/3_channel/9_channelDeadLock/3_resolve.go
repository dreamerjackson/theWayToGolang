/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
)
//死锁的解决方法2:
func doWork(id int,
	c chan int,done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		go func() { done <- true}()
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

func chanDemo() {


	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}


	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _,worker := range workers{
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()
}