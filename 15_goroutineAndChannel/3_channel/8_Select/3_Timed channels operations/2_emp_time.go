/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second) //10秒后往通道发送消息
	tick := time.Tick(time.Second) //每隔一秒钟发送一次
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: //通道为nil时不会传递任何值
			values = values[1:]

		case <-time.After(800 * time.Millisecond)://代表case会在800毫秒传递值到通道，如果在800毫秒之内其他的case有反应，就不会执行，但是如果其他case在800毫秒没有反应就会执行此case
			fmt.Println("timeout")
		case <-tick:   //每隔一秒往通道发送数据
			fmt.Println(
				"queue len =", len(values))
		case <-tm:// time.After(10 * time.Second) 放在外部，10秒后发送数据
			fmt.Println("bye")
			return
		}
	}
}