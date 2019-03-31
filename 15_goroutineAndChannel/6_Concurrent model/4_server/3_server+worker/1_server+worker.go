/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

//
//
//Server with worker example is a bit advanced version of the logger.
//	It not only does some work but sends the result of its work back to the pool using results channel.
//	Not a big deal, but it extends our logger example to something more practical.
//Let’s see the code and animation:


package main

import (
	"net"
	"time"
)

func handler(c net.Conn, ch chan string) {
	addr := c.RemoteAddr().String()
	ch <- addr
	time.Sleep(100 * time.Millisecond)
	c.Write([]byte("ok"))
	c.Close()
}

func logger(wch chan int, results chan int) {
	for {
		data := <-wch
		data++
		results <- data
	}
}

func parse(results chan int) {
	for {
		<-results
	}
}

func pool(ch chan string, n int) {
	wch := make(chan int)
	results := make(chan int)
	for i := 0; i < n; i++ {
		go logger(wch, results)
	}
	go parse(results)
	for {
		addr := <-ch
		l := len(addr)
		wch <- l
	}
}

func server(l net.Listener, ch chan string) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, ch)
	}
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	ch := make(chan string)
	go pool(ch, 4)
	go server(l, ch)
	time.Sleep(10 * time.Second)
}