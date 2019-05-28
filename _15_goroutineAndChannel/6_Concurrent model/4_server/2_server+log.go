


/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */


/*
It’s not very interesting - it seems there is nothing happens in terms of concurrency.
Of course, under the hood there is a ton of complexity, which is deliberately hidden from us. “Simplicity is complicated”.

But let’s go back to concurrency and add some interaction to our server.
Let’s say, each handler wants to write asynchronously to the logger. Logger itself,
in our example, is a separate goroutine which does the job.

*/
package main

import (
	"fmt"
	"net"
	"time"
)

func handler(c net.Conn, ch chan string) {
	ch <- c.RemoteAddr().String()
	c.Write([]byte("ok"))
	c.Close()
}

func logger(ch chan string) {
	for {
		fmt.Println(<-ch)
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
	go logger(ch)
	go server(l, ch)
	time.Sleep(10 * time.Second)
}
