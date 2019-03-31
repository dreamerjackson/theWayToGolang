/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"time"
)

//通道作为一等公民，作为参数和返回值
func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}


func main(){

	chanDemo()
}

