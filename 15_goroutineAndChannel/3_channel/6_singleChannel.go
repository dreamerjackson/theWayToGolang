/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

/*
当一个channel作为一个函数参数是，它一般总是被专门用于只发送或者只接收。
为了表明这种意图并防止被滥用，Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel。
类型chan<- int表示一个只发送int的channel，只能发送不能接收。
相反，类型<-chan int表示一个只接收int的channel，只能接收不能发送。
（箭头<-和关键字chan的相对位置表明了channel的方向。）这种限制将在编译期检测。
因为关闭操作只用于断言不再向channel发送新的数据，
所以只有在发送者所在的goroutine才会调用close函数，
因此对一个只接收的channel调用close将是一个编译错误。

*/
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

/*
调用counter(naturals)将导致将chan int类型的naturals隐式地转换为chan<- int类型只发送型的channel。
调用printer(squares)也会导致相似的隐式转换，这一次是转换为<-chan int类型只接收型的channel。
任何双向channel向单向channel变量的赋值操作都将导致该隐式转换。
这里并没有反向转换的语法：也就是不能一个将类似chan<- int类型的单向型的channel转换为chan int类型的双向型的channel。

*/