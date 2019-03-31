/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"net"
	"io"
	"os"
	"log"
)

//不带缓存的Channels 可以作为锁使用

//一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，
// 直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
// 反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作。
//基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。
// 因为这个原因，无缓存Channels有时候也被称为同步Channels。
// 当通过一个无缓存Channels发送数据时，接收者收到数据发生在唤醒发送者goroutine之前（译注：happens before，这是Go语言并发内存模型的一个关键术语！）。
//在讨论并发编程时，当我们说x事件在y事件之前发生（happens before），
// 我们并不是说x事件在时间上比y时间更早；我们要表达的意思是要保证在此之前的事件都已经完成了，
// 例如在此之前的更新某些变量的操作已经完成，你可以放心依赖这些已完成的事件了。
//当我们说x事件旣不是在y事件之前发生也不是在y事件之后发生，我们就说x事件和y事件是并发的。
// 这并不是意味着x事件和y事件就一定是同时发生的，我们只是不能确定这两个事件发生的先后顺序。
// 在后面我们将看到，当两个goroutine并发访问了相同的变量时，我们有必要保证某些事件的执行顺序，以避免出现某些并发问题。


func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}