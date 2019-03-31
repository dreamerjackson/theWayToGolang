/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main


// NOTE: the ticker goroutine never terminates if the launch is aborted.
// This is a "goroutine leak".


//如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会。
// 增加前一个例子的buffer大小会使其输出变得不确定，因为当buffer旣不为满也不为空时，select语句的执行情况就像是抛硬币的行为一样是随机的。
//下面让我们的发射程序打印倒计时。这里的select语句会使每次循环迭代等待一秒来执行退出操作。
import (
	"fmt"
	"os"
	"time"
)

//!+

func main() {
	// ...create abort channel...

	//!-
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}

/*
time.Tick函数表现得好像它创建了一个在循环中调用time.Sleep的goroutine，每次被唤醒时发送一个事件。
当countdown函数返回时，它会停止从tick中接收事件，但是ticker这个goroutine还依然存活，
继续徒劳地尝试从channel中发送值，然而这时候已经没有其它的goroutine会从该channel中接收值了--这被称为goroutine泄露(§8.4.4)。
Tick函数挺方便，但是只有当程序整个生命周期都需要这个时间时我们使用它才比较合适。否则的话，我们应该使用下面的这种模式：
ticker := time.NewTicker(1 * time.Second)
<-ticker.C    // receive from the ticker's channel
ticker.Stop() // cause the ticker's goroutine to terminate
有时候我们希望能够从channel中发送或者接收值，并避免因为发送或者接收导致的阻塞，尤其是当channel没有准备好写或者读时。
select语句就可以实现这样的功能。select会有一个default来设置当其它的操作都不能够马上被处理时程序需要执行哪些逻辑。
下面的select语句会在abort channel中有值时，从其中接收值；无值时什么都不做。这是一个非阻塞的接收操作；反复地做这样的操作叫做“轮询channel”。
```
select {
case <-abort:
    fmt.Printf("Launch aborted!\n")
    return
default:
    // do nothing
}
```
channel的零值是nil。也许会让你觉得比较奇怪，nil的channel有时候也是有一些用处的。
因为对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。
这使得我们可以用nil来激活或者禁用case，来达成处理其它输入或输出事件时超时和取消的逻辑。
*/