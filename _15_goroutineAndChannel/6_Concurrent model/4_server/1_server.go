


/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */


//
//Next common pattern is similar to fan-out, but with goroutines spawned for the short period of time, just to accomplish some task.
//	It’s typically used for implementing servers - create a listener, run accept() in a loop and start goroutine for each accepted connection.
//	It’s very expressive and allows to implement server handlers as simple as possible. Take a look at this simple example:


package main

import "net"

func handler(c net.Conn) {
	c.Write([]byte("ok"))
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
