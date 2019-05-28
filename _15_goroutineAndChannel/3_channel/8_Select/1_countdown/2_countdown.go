/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"os"
	"time"
)

//现在我们让这个程序支持在倒计时中，用户按下return键时直接中断发射流程。
//首先，我们启动一个goroutine，这个goroutine会尝试从标准输入中调入一个单独的byte并且，如果成功了，会向名为abort的channel发送一个值。
/*
现在每一次计数循环的迭代都需要等待两个channel中的其中一个返回事件了：ticker channel当一切正常时(就像NASA jorgon的"nominal"，译注：这梗估计我们是不懂了)或者异常时返回的abort事件。我们无法做到从每一个channel中接收信息，如果我们这么做的话，如果第一个channel中没有事件发过来那么程序就会立刻被阻塞，这样我们就无法收到第二个channel中发过来的事件。这时候我们需要多路复用(multiplex)这些操作了，为了能够多路复用，我们使用了select语句。
select {
case <-ch1:
    // ...
case x := <-ch2:
    // ...use x...
case ch3 <- y:
    // ...
default:
    // ...
}
上面是select语句的一般形式。和switch语句稍微有点相似，也会有几个case和最后的default选择支。
每一个case代表一个通信操作(在某个channel上进行发送或者接收)并且会包含一些语句组成的一个语句块。
一个接收表达式可能只包含接收表达式自身(译注：不把接收到的值赋值给变量什么的)，就像上面的第一个case，或者包含在一个简短的变量声明中，
像第二个case里一样；第二种形式让你能够引用接收到的值。
select会等待case中有能够执行的case时去执行。当条件满足时，select才会去通信并执行case之后的语句；
这时候其它通信是不会执行的。一个没有任何case的select语句写作select{}，会永远地等待下去。
让我们回到我们的火箭发射程序。time.After函数会立卽返回一个channel，并起一个新的goroutine在经过特定的时间后向该channel发送一个独立的值。
下面的select语句会会一直等待到两个事件中的一个到达，无论是abort事件或者一个10秒经过的事件。如果10秒经过了还没有abort事件进入，那么火箭就会发射。

*/
func main() {
	// ...create abort channel...

	//!-

	//!+abort
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	//!-abort

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}

