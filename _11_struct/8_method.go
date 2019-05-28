/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"math"
	"fmt"
)



//在函数声明时，在其名字之前放上一个变量，即是一个方法。
//这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。
type Point struct{ X, Y float64 }

//  传统函数 traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 方法 same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//上面的代码里那个附加的参数p，叫做方法的接收器(receiver)，
//早期的面向对象语言留下的遗产将调用一个方法称为“向一个对象发送消息”。
//
//在Go语言中，我们并不会像其它语言那样用this或者self作为接收器；
//我们可以任意的选择接收器的名字。
//由于接收器的名字经常会被使用到，所以保持其在方法间传递时的一致性和简短性是不错的主意。
//这里的建议是可以使用其类型的第一个字母，比如这里使用了Point的首字母p。


func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}


//对于一个给定的类型，其内部的方法都必须有唯一的方法名，
//但是不同的类型却可以有同样的方法名，比如我们这里Point和Path就都有Distance这个名字的方法；
//所以我们没有必要非在方法名之前加类型名来消除歧义，比如PathDistance。
//这里我们已经看到了方法比之函数的一些好处：方法名可以简短。
//当我们在包外调用的时候这种好处就会被放大，因为我们可以使用这个短名字，而可以省略掉包的名字
type Path []Point
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main(){

	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call


	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"



	//------方法值---------
	p = Point{1, 2}
	q = Point{4, 6}

	distanceFromP := p.Distance        // method value
	fmt.Println(distanceFromP(q))      // "5"

	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)

	scaleP := p.ScaleBy // method value
	scaleP(2)           // p becomes (2, 4)
	scaleP(3)           //      then (6, 12)
	scaleP(10)          //      then (60, 120)

	//------方法表达式 第一个参数需要指定调用者---------
	//当T是一个类型时，方法表达式可能会写作T.f或者(*T).f，会返回一个函数"值"，
	//这种函数会将其第一个参数用作接收器
	p = Point{1, 2}
	q = Point{4, 6}

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
}