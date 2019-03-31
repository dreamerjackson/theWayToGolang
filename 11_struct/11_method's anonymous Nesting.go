/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"math"
)

//!+decl
import (
	"image/color"
	"sync"
)

type Point struct{ X, Y float64 }
//匿名包含方法 我们可以把ColoredPoint类型当作接收器来调用Point里的方法，即使ColoredPoint里没有声明这些方法
//Point类的方法也被引入了ColoredPoint
type ColoredPoint struct {
	Point
	Color color.RGBA
}

//!-decl

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}



func main() {
	//!+main
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
	//!-main
}


/*
//一个ColoredPoint并不是一个Point，但他"has a"Point，并且它有从Point类里引入的Distance和ScaleBy方法。
//如果你喜欢从实现的角度来考虑问题，内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法，和下面的形式是等价的：

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}
*/


//小tip  将锁放置在结构体中，就可以使用锁的方法。
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}


func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}