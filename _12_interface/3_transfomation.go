


/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */


//接口对象如何转换为实际类型
package main

import (
	"math"
	"fmt"
)

//1、定义接口
type Shape interface {
	perimeter() float64
	area() float64
}

//2.矩形
type Rectangle struct {
	a, b float64
}

//3.三角形
type Triangle struct {
	a, b, c float64
}

//4.圆形
type Circle struct {
	radius float64
}

//定义实现接口的方法
func (r Rectangle) perimeter() float64 {
	return (r.a + r.b) * 2
}

func (r Rectangle) area() float64 {
	return r.a * r.b
}

func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	//海伦公式
	p := t.perimeter() / 2 //半周长
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

//接口对象转型方式1
//instance,ok := 接口对象.(实际类型)
func getType(s Shape) {
	if instance, ok := s.(Rectangle); ok {
		fmt.Printf("矩形：长度%.2f , 宽度%.2f , ", instance.a, instance.b)
	} else if instance, ok := s.(Triangle); ok {
		fmt.Printf("三角形：三边分别：%.2f , %.2f , %.2f , ", instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Circle); ok {
		fmt.Printf("圆形：半径%.2f , ", instance.radius)
	}
}

//接口对象转型——方式2
//接口对象.(type),  配合switch和case语句使用
func getType2(s Shape) {
	switch instance := s.(type) {
	case Rectangle:
		fmt.Printf("矩形：长度为%.2f ， 宽为%.2f ，\t", instance.a, instance.b)
	case Triangle:
		fmt.Printf("三角形：三边分别为%.2f ，%.2f ， %.2f ，\t", instance.a, instance.b, instance.c)
	case Circle:
		fmt.Printf("圆形：半径为%.2f ，\t", instance.radius)
	}
}

func getResult(s Shape) {
	getType2(s)
	fmt.Printf("周长：%.2f ，面积:%.2f \n", s.perimeter(), s.area())
}

func main() {
	var s Shape
	s = Rectangle{3, 4}
	getResult(s)
	showInfo(s)

	s = Triangle{3, 4, 5}
	getResult(s)
	showInfo(s)

	s = Circle{1}
	getResult(s)
	showInfo(s)

	x := Triangle{3, 4, 5}
	fmt.Println(x)

}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle对象，属性分别为：%.2f, %.2f, %.2f", t.a, t.b, t.c)
}

func showInfo(s Shape) {
	fmt.Printf("%T ,%v \n", s, s)
	fmt.Println("-------------------")
}