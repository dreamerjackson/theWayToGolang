/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

/*
如您所料，上面的程序将生成0、1、4、9、……形式的无穷数列。
像这样的串联Channels的管道（Pipelines）可以用在需要长时间运行的服务中，每个长时间运行的goroutine可能会包含一个死循环，
在不同goroutine的死循环内部使用串联的Channels来通信。
但是，如果我们希望通过Channels只发送有限的数列该如何处理呢？
如果发送者知道，没有更多的值需要发送到channel的话，那么让接收者也能及时知道没有多余的值可接收将是有用的，
因为接收者可以停止不必要的接收等待。这可以通过内置的close函数来关闭channel实现：
close(naturals)
当一个channel被关闭后，再向该channel发送数据将导致panic异常。
当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立卽返回一个零值。
关闭上面例子中的naturals变量对应的channel并不能终止循环，它依然会收到一个永无休止的零值序列，然后将它们发送给打印者goroutine。
没有办法直接测试一个channel是否被关闭，但是接收操作有一个变体形式：
它多接收一个结果，多接收的第二个结果是一个布尔值ok，ture表示成功从channels接收到值，false表示channels已经被关闭并且里面没有值可接收。
使用这个特性，我们可以修改squarer函数中的循环代码，当naturals对应的channel被关闭并没有值可接收时跳出循环，并且也关闭squares对应的channel.
*/

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0;x<100 ; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}


