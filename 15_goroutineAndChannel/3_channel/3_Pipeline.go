/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

/*
  Channels也可以用于将多个goroutine链接在一起，一个Channels的输出作为下一个Channels的输入。
这种串联的Channels就是所谓的管道（pipeline）。下面的程序用两个channels将三个goroutine串联起来
第一个goroutine是一个计数器，用于生成0、1、2、……形式的整数序列，
然后通过channel将该整数序列发送给第二个goroutine；第二个goroutine是一个求平方的程序，对收到的每个整数求平方，
然后将平方后的结果通过第二个channel发送给第三个goroutine；
第三个goroutine是一个打印程序，打印收到的每个整数。
为了保持例子清晰，我们有意选择了非常简单的函数，当然三个goroutine的计算很简单，在现实中确实没有必要为如此​​简单的运算构建三个goroutine。

*/
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

