

/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

//扇入 多个入，抢夺一个通道出

package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 100*time.Millisecond)
	go producer(ch, 250*time.Millisecond)
	go reader(out)
	for i := range ch {
		out <- i
	}
}