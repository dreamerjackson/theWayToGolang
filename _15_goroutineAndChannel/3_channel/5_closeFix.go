/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

/*
因为上面的语法是笨拙的，而且这种处理模式很常见，因此Go语言的range循环可直接在channels上面迭代。
使用range循环是上面处理模式的简洁语法，它依次从channel接收数据，当channel被关闭并且没有值可接收时跳出循环
在下面的改进中，我们的计数器goroutine只生成100个含数字的序列，然后关闭naturals对应的channel，
这将导致计算平方数的squarer对应的goroutine可以正常终止循环并关闭squares对应的channel。
（在一个更复杂的程序中，可以通过defer语句关闭对应的channel。最后，主goroutine也可以正常终止循环并退出程序。
*/

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

/*
其实你并不需要关闭每一个channel。
只要当需要告诉接收者goroutine，所有的数据已经全部发送时才需要关闭channel。
不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。
（不要将关闭一个打开文件的操作和关闭一个channel操作混淆。对于每个打开的文件，都需要在不使用的使用调用对应的Close方法来关闭文件。）
试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。关闭一个channels还会触发一个广播机制

*/