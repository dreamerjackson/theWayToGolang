/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；
//这类成员就叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。
//下面的代码中，Circle和Wheel各自都有一个匿名成员。我们可以说Point类型被嵌入到了Circle结构体，同时Circle类型被嵌入到了Wheel结构体。

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main(){
	var w Wheel
	w.X = 8            // equivalent to w.Circle.Point.X = 8
	w.Y = 8            // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5       // equivalent to w.Circle.Radius = 5
	w.Spokes = 20
	/*
	在右边的注释中给出的显式形式访问这些叶子成员的语法依然有效，因此匿名成员并不是眞的无法访问了。其中匿名成员Circle和Point都有自己的名字——就是命名的类型名字——但是这些名字在点操作符中是可选的。我们在访问子成员的时候可以忽略任何匿名成员部分。
	不幸的是，结构体字面值并没有简短表示匿名成员的语法， 因此下面的语句都不能编译通过：

	//w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	//w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
	*/

	//结构体字面值必须遵循形状类型声明时的结构，所以我们只能用下面的两种语法，它们彼此是等价的：
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)


	w.X = 42

	fmt.Printf("%#v\n", w)
}


//需要注意的是Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字。
//因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。
//同时，因为成员的名字是由其类型隐式地决定的，所有匿名成员也有可见性的规则约束。
//在上面的例子中，Point和Circle匿名成员都是导出的。卽使它们不导出（比如改成小写字母开头的point和circle），我们依然可以用简短形式访问匿名成员嵌套的成员
//w.X = 8 // equivalent to w.circle.point.X = 8
//但是在包外部，因为circle和point没有导出不能访问它们的成员，因此简短的匿名成员访问语法也是禁止的。


//到目前未知，我们看到匿名成员特性只是对访问嵌套成员的点运算符提供了简短的语法糖。稍后，我们将会看到匿名成员并不要求是结构体类型；
// 其实任何命令的类型都可以作为结构体的匿名成员。但是为什么要嵌入一个没有任何子成员类型的匿名成员类型呢？
//答案是匿名类型的方法集。简短的点运算符语法可以用于选择匿名成员嵌套的成员，也可以用于访问它们的方法。
// 实际上，外层的结构体不仅仅是获得了匿名成员类型的所有成员，而且也获得了该类型导出的全部的方法。这个机制可以用于将一个有简单行为的对象组合成有复杂行为的对象