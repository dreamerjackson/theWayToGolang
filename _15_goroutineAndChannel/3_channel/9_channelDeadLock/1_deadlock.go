/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main
//通道的经典死锁现象
import (
	"fmt"
)

func doWork(id int,
	c chan int,done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		//done 代表任务做完，需要人来接收。因为done一直等不到人来接收，就会卡住，卡住就不会接收C
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

func chanDemo() {


	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	//第一遍循环顺利进行
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	//第二遍循环会陷入死锁，原因见上。
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
