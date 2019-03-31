/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//当调用一个函数时，会对其每一个参数值进行拷贝，
//如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，这种情况下我们就需要用到指针了。
//对应到我们这里用来更新接收器的对象的方法，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法，如下：
type Point struct{ X, Y float64 }
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//想要调用指针类型方法(*Point).ScaleBy，只要提供一个Point类型的指针即可，像下面这样
func main(){
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"

	//如果接收器p是一个Point类型的变量，并且其方法需要一个Point指针作为接收器，
	// 我们可以用下面这种简短的写法
	p := Point{1, 2}
	p.ScaleBy(2)
}

//不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
//在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，
//第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；
//第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址